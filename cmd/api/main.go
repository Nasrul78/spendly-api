//	@title			Spendly API
//	@version		1.0
//	@description	A personal expense tracking REST API.

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description				Type Bearer followed by a space and your JWT token.

package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/nasrul78/spendly-api/config"
	_ "github.com/nasrul78/spendly-api/docs"
	"github.com/nasrul78/spendly-api/internal/handler"
	"github.com/nasrul78/spendly-api/internal/middleware"
	"github.com/nasrul78/spendly-api/internal/repository"
	"github.com/nasrul78/spendly-api/internal/service"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg := config.Load()

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		slog.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer pool.Close()

	if err := pool.Ping(context.Background()); err != nil {
		slog.Error("failed to reach database", "error", err)
		os.Exit(1)
	}
	slog.Info("successfully connected to database")

	// repositories
	userRepo := repository.NewUserRepository(pool)
	categoryRepo := repository.NewCategoryRepository(pool)
	expenseRepo := repository.NewExpenseRepository(pool)

	// services
	authService := service.NewAuthService(userRepo, cfg)
	categoryService := service.NewCategoryService(categoryRepo)
	expenseService := service.NewExpenseService(expenseRepo)

	// handlers
	authHandler := handler.NewAuthHandler(authService)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	expenseHandler := handler.NewExpenseHandler(expenseService)

	// router
	r := chi.NewRouter()
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/register", authHandler.Register)
			r.Post("/login", authHandler.Login)
		})

		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(cfg.JWTSecret))

			r.Route("/categories", func(r chi.Router) {
				r.Post("/", categoryHandler.Create)
				r.Get("/", categoryHandler.GetAll)
				r.Get("/{id}", categoryHandler.GetByID)
				r.Put("/{id}", categoryHandler.Update)
				r.Delete("/{id}", categoryHandler.Delete)
			})

			r.Route("/expenses", func(r chi.Router) {
				r.Get("/summary", expenseHandler.GetSummaryByUserID)
				r.Post("/", expenseHandler.Create)
				r.Get("/", expenseHandler.GetAll)
				r.Get("/{id}", expenseHandler.GetByID)
				r.Put("/{id}", expenseHandler.Update)
				r.Delete("/{id}", expenseHandler.Delete)
			})
		})
	})

	addr := ":" + cfg.AppPort
	slog.Info("starting server", "address", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		slog.Error("failed to start server", "error", err)
		os.Exit(1)
	}
}
