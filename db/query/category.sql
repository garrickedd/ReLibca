-- name: AddCategory :one
INSERT INTO Categories (Id, Name, Type)
VALUES (DEFAULT, $1, $2)
RETURNING *;