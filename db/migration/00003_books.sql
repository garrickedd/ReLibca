-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
create table
    Books (
        Id uuid primary key not null default uuid_generate_v4(),
        Name text not null,
        Author text,
        Place text,
        CreatedAt timestamp not null default current_timestamp,
        ModifiedAt timestamp default current_timestamp,
        CategoryId UUID REFERENCES Categories(Id)
    )
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists Books;
-- +goose StatementEnd
