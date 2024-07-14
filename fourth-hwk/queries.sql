-- name: CreateBankAccount :exec
INSERT INTO users (username, balance) VALUES ($1, $2);

-- name: GetBankAccount :one
SELECT * FROM users WHERE username = $1;

-- name: UpdateBankAccountBalance :exec
UPDATE users SET balance = $2 WHERE username = $1;

-- name: UpdateBankAccountName :exec
UPDATE users SET username = $2 WHERE username = $1;


-- name: DeleteBankAccount :exec
DELETE FROM users WHERE username = $1;

-- name: GetAllBankAccounts :many
SELECT * FROM users;