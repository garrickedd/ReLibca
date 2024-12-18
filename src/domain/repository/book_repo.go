package repository

import (
	"errors"
	"fmt"
	"log"

	"github.com/garrickedd/ReLibca/src/domain/model"
	"github.com/garrickedd/ReLibca/src/shared/config"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
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
		place
	) 
		VALUES(
			:title, 
			:author, 
			:count, 
			:tag, 
			:place
		)`

	// Convert the tag slice to pq.Array
	// data.Tag = pq.Array(data.Tag)

	_, err := r.NamedExec(queryBook, map[string]interface{}{
		"title":  data.Title,
		"author": data.Author,
		"count":  data.Count,
		"tag":    pq.Array(data.Tag), // Pass pq.Array directly here
		"place":  data.Place,
	})
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
	place = :place
	WHERE title = :title
	`

	_, err := r.NamedExec(queryBook, map[string]interface{}{
		"title":  data.Title,
		"author": data.Author,
		"count":  data.Count,
		"tag":    pq.Array(data.Tag), // Pass pq.Array directly here
		"place":  data.Place,
	})
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	return &config.Result{Message: "1 book updated"}, nil
}

// TODO: Fix "sql: Scan error on column index 4, name \"tag\": unsupported Scan, storing driver.Value type []uint8 into type *[]string"
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

	// Calculate total record count for pagination
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

	// Query to search by title only
	if search != "" {
		// Use ILIKE for case-insensitive search in title
		query := `SELECT * FROM books WHERE title ILIKE $1 LIMIT $2 OFFSET $3`
		err := r.Select(&books_data, query, "%"+search+"%", limit, offset)
		if err != nil {
			return nil, nil, err
		}
	} else if orderby != "" {
		// Search and order by specified column
		query := `SELECT * FROM books ORDER BY ` + orderby + ` ASC LIMIT $1 OFFSET $2`
		err := r.Select(&books_data, query, limit, offset)
		if err != nil {
			return nil, nil, err
		}
	} else if search != "" && orderby != "" {
		// Combine title search and ordering
		query := `SELECT * FROM books WHERE title ILIKE $1 ORDER BY ` + orderby + ` ASC LIMIT $2 OFFSET $3`
		err := r.Select(&books_data, query, "%"+search+"%", limit, offset)
		if err != nil {
			return nil, nil, err
		}
	} else {
		// Simple query without filtering
		query := `SELECT * FROM books LIMIT $1 OFFSET $2`
		err := r.Select(&books_data, query, limit, offset)
		if err != nil {
			return nil, nil, err
		}
	}

	// Ensure that the "tag" field is correctly scanned as a string array
	for i := range books_data {
		// Use pq.Array to scan the tag field
		if err := pq.Array(&books_data[i].Tag).Scan(books_data[i].Tag); err != nil {
			return nil, nil, err
		}
	}

	// If no books are found
	if len(books_data) == 0 {
		return nil, nil, errors.New("data not found")
	}

	return books_data, pgnt, nil
}
