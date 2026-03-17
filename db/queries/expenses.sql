-- name: CreateExpense :one
INSERT INTO expenses (user_id, category_id, amount, note, date)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetExpenseByID :one
SELECT e.*, c.name as category_name
FROM expenses e
LEFT JOIN categories c ON e.category_id = c.id
WHERE e.id = $1 AND e.user_id = $2;

-- name: GetExpensesByUserID :many
SELECT e.*, c.name as category_name
FROM expenses e
LEFT JOIN categories c ON e.category_id = c.id
WHERE e.user_id = $1
  AND ($2::date IS NULL OR e.date >= $2::date)
  AND ($3::date IS NULL OR e.date <= $3::date)
  AND ($4::uuid IS NULL OR e.category_id = $4::uuid)
ORDER BY e.date DESC
LIMIT $5 OFFSET $6;

-- name: CountExpensesByUserID :one
SELECT COUNT(*) FROM expenses
WHERE user_id = $1
  AND ($2::date IS NULL OR date >= $2::date)
  AND ($3::date IS NULL OR date <= $3::date)
  AND ($4::uuid IS NULL OR category_id = $4::uuid);

-- name: UpdateExpense :one
UPDATE expenses
SET category_id = $1, amount = $2, note = $3, date = $4, updated_at = NOW()
WHERE id = $5 AND user_id = $6
RETURNING *;

-- name: DeleteExpense :exec
DELETE FROM expenses
WHERE id = $1 AND user_id = $2;

-- name: GetExpenseSummaryByUserID :many
SELECT c.name as category_name, SUM(e.amount) as total_amount
FROM expenses e
LEFT JOIN categories c ON e.category_id = c.id
WHERE e.user_id = $1
  AND ($2::date IS NULL OR e.date >= $2::date)
  AND ($3::date IS NULL OR e.date <= $3::date)
GROUP BY c.name
ORDER BY total_amount DESC;
