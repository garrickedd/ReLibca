-- name: CreateUser :one
insert into users(
    FullName,
    UserName,
    Password,
    PhoneNumber,
    Role,
    IsActive,
    CreatedAt,
    ModifiedAt
) values (
    $1, $2, $3, $4, $5, $6, $7, $8
) returning *;

-- name: GetUserById :one
select * from users
where Id = $1 LIMIT 1;

-- name: ListUsers :many
select * from users
order by Id
limit $1
offset $2;

-- name: UpdateUser :one
update users
set
FullName = coalesce(sqlc.narg('FullName'), FullName),
Password = coalesce(sqlc.narg('Password'), Password),
PhoneNumber = coalesce(sqlc.narg('PhoneNumber'), PhoneNumber)
where Id = sqlc.arg('Id')
returning *;

-- name: SetActiveStatus :one
update users
set IsActive = $2, ModifiedAt = current_timestamp
where Id = $1
returning *;

-- name: DeleteUser :exec
delete from users
where Id = $1;