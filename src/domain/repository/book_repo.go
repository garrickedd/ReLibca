package repository

import (
	"errors"
	"fmt"
	"log"

	"github.com/garrickedd/ReLibca/src/domain/model"
	"github.com/garrickedd/ReLibca/src/shared/config"
	"github.com/jmoiron/sqlx"
)

type RepoBookIF interface {
	CreateBook(data *model.Book) (*config.Result, error)
	UpdateBook(data *model.Book) (*config.Result, error)
	DeleteBook(data *model.Book) (*config.Result, error)
	Get_Books(data *model.Book, search string, limit int, page int, orderby string) ([]model.Book, *Pagination, error)
}

type RepoBook struct {
	*sqlx.DB
}

func NewBook(db *sqlx.DB) *RepoBook {
	return &RepoBook{db}
}

func (r RepoBook) CreateBook(data *model.Book) (*config.Result, error) {
	queryBook := `INSERT INTO books(
		title,  
		author, 
		count,         
		tag,         
		place,
		image_file
	) 
		VALUES(
			:title, 
			:author, 
			:count, 
			:tag, 
			:place,
			:image_file
		)`

	_, err := r.NamedExec(queryBook, data)
	if err != nil {
		return nil, err
	}
	return &config.Result{Message: "1 book created"}, nil
}

func (r RepoBook) DeleteBook(data *model.Book) (*config.Result, error) {
	queryBook := `DELETE FROM public.books WHERE id_book=:id_book`
	_, err := r.NamedExec(queryBook, data)
	if err != nil {
		return nil, err
	}
	return &config.Result{Message: "1 book deleted"}, nil
}

func (r RepoBook) UpdateBook(data *model.Book) (*config.Result, error) {
	queryBook := `UPDATE public.books SET
	title = :title,
	author = :author,
	count = :count,
	tag = :tag,
	place = :place,
	image_file = :image_file,
	WHERE id_book = :id_book
	`

	_, err := r.NamedExec(queryBook, data)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	return &config.Result{Message: "1 book updated"}, nil
}

func (r RepoBook) Get_Books(data *model.Book, search string, limit int, page int, orderby string) ([]model.Book, *Pagination, error) {
	books_data := []model.Book{}

	var (
		pgnt        = &Pagination{}
		recordcount int
	)

	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 5
	}

	offset := limit * (page - 1)

	sqltable := fmt.Sprintln("SELECT count(id_book) FROM books")

	r.QueryRow(sqltable).Scan(&recordcount)

	total := (recordcount / limit)

	remainder := (recordcount % limit)
	if remainder == 0 {
		pgnt.TotalPage = total
	} else {
		pgnt.TotalPage = total + 1
	}

	pgnt.CurrentPage = page
	pgnt.RecordPerPage = limit

	if page <= 0 {
		pgnt.Next = page + 1
	} else if page < pgnt.TotalPage {
		pgnt.Previous = page - 1
		pgnt.Next = page + 1
	} else if page == pgnt.TotalPage {
		pgnt.Previous = page - 1
		pgnt.Next = 0
	}

	if search != "" {
		r.Select(&books_data, `select * from books WHERE title=$1 LIMIT $2 OFFSET $3 `, search, limit, offset)
	} else if orderby != "" {
		r.Select(&books_data, `select * from books order by `+orderby+` asc LIMIT $1 OFFSET $2`, limit, offset)
	} else if orderby != "" && search != "" {
		r.Select(&books_data, `select * from books order by `+orderby+` asc WHERE title=$1 LIMIT $2 OFFSET $3 `, search, limit, offset)
	} else {
		r.Select(&books_data, `select * from books LIMIT $1 OFFSET $2 `, limit, offset)
	}
	if len(books_data) == 0 {
		return nil, nil, errors.New("data not found")
	}
	return books_data, pgnt, nil
}
