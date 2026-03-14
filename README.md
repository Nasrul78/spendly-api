# Spendly API

> A personal expense tracking REST API built with Go.

Spendly is a RESTful API built with Go for tracking personal expenses and financial habits. Users can register, authenticate via JWT, and manage their own categories and expense records — all scoped privately per account. The API supports filtering expenses by date range and category, as well as generating spending summaries to give users a clear picture of where their money goes.

Designed with clean architecture principles, Spendly separates concerns across handler, service, and repository layers, using PostgreSQL as its database and `sqlc` for type-safe query generation.

---

## Tech Stack

| Layer | Tool |
|---|---|
| Language | Go 1.22+ |
| Router | `chi` |
| Database | PostgreSQL |
| Query Generation | `sqlc` |
| Migrations | `goose` |
| Validation | `go-playground/validator` |
| Auth | `golang-jwt/jwt` |
| Config | `godotenv` |
| Dev Reload | `air` |
| Testing | `testify` + `httptest` |
| Container | Docker + docker-compose |

---

## Project Structure

```
spendly/
├── cmd/
│   └── api/
│       └── main.go
├── .air.toml
├── .env
├── .env.example
├── .gitignore
├── go.mod
├── go.sum
└── README.md