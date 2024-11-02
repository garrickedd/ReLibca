-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
create table
    Users (
        Id uuid primary key not null default uuid_generate_v4(),
        FullName text not null,
        UserName varchar(256) not null,
        Password varchar(256) not null,
        PhoneNumber varchar(12) not null,
        Role int not null,
        IsActive boolean not null default true,
        CreatedAt timestamp not null default current_timestamp,
        ModifiedAt timestamp default current_timestamp
    )
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
drop table if exists Users;
-- +goose StatementEnd