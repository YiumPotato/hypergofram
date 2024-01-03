-- name: CreateTodo :one
INSERT INTO todos (
    title varchar,
    done bool
)
VALUES ($1, $2)
RETURNING *;

-- name: GetTodo :one
SELECT * FROM todos
WHERE id = $1;

-- name: ListTodos :many
SELECT * FROM todos
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateTodo :one
UPDATE todos
SET title = $1,
WHERE id = $2
RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = $1;

-- name: CountTodos :one
SELECT COUNT(*) FROM todos;

-- name: ClearTodos :exec
DELETE FROM todos;
