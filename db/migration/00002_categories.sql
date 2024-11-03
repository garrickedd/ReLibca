-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE Categories (
	Id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	Name TEXT NOT NULL,
	Type INT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists Categories;
-- +goose StatementEnd
