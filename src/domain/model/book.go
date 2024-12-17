package model

// id_book uuid NOT NULL DEFAULT uuid_generate_v4(),
// title varchar(255) NOT NULL,
// author varchar(255) NOT NULL,
// count int NOT NULL DEFAULT 0,
// tag text[] NOT NULL,
// place varchar(255),

type Book struct {
	Id_book string   `db:"id_book" form:"id_book" json:"id_book" `
	Title   string   `db:"title" form:"title" json:"title" `
	Author  string   `db:"author" form:"author" json:"author" `
	Count   int      `db:"count" form:"count" json:"count" `
	Tag     []string `db:"tag" form:"tag" json:"tag" `
	Place   string   `db:"place" form:"place" json:"place" `
}
