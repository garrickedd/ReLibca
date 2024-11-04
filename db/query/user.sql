-- name: CreateUser :one
insert into Users(
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
select * from Users
where Id = $1 LIMIT 1;

-- name: ListUsers :many
select * from Users
order by Id
limit $1
offset $2;

-- name: UpdateUser :one
update Users
set
FullName = coalesce(sqlc.narg('FullName'), FullName),
Password = coalesce(sqlc.narg('Password'), Password),
PhoneNumber = coalesce(sqlc.narg('PhoneNumber'), PhoneNumber)
where Id = sqlc.arg('Id')
returning *;

-- name: SetActiveStatus :one
update Users
set IsActive = $2, ModifiedAt = current_timestamp
where Id = $1
returning *;

-- name: DeleteUser :exec
delete from Users
where Id = $1;