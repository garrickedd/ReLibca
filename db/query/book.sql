-- name: AddBook :one
insert into Books(
    Id,
    Name,
    Author,
    Place,
    CreatedAt,
    ModifiedAt,
    CategoryId
) values (
    default, $1, $2, $3, default, null, $4
) returning *;

-- name: ListBooks :many
select 
    Books.Id,
    Books.Name,
    Books.Author,
    Books.Place,
    Categories.Name as CategoryName
from
    Books
left join
    Categories on Books.CategoryId = Categories.Id
where
    $1 is null or Books.Name ilike '%' || $1 || '%'
order by Books.CreatedAt desc;

-- name: UpdateBook :one

