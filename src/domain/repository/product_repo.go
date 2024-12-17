package repository

import (
	"errors"
	"fmt"
	"log"

	"github.com/garrickedd/ReLibca/src/domain/model"
	"github.com/garrickedd/ReLibca/src/shared/config"
	"github.com/jmoiron/sqlx"
)

type RepoProductIF interface {
	CreateProduct(data *model.Product) (*config.Result, error)
	UpdateProduct(data *model.Product) (*config.Result, error)
	DeleteProduct(data *model.Product) (*config.Result, error)
	Get_Products(data *model.Product, search string, limit int, page int, orderby string) ([]model.Product, *Pagination, error)
}

type RepoProduct struct {
	*sqlx.DB
}

type Pagination struct {
	Next          int
	Previous      int
	RecordPerPage int
	CurrentPage   int
	TotalPage     int
}

func NewProduct(db *sqlx.DB) *RepoProduct {
	return &RepoProduct{db}
}

func (r RepoProduct) CreateProduct(data *model.Product) (*config.Result, error) {
	queryproduct := `INSERT INTO products(
		product_name,  
		description, 
		status,         
		price,         
		image_file,
		categories,
		isfavorite
		) 
		VALUES(
			:product_name, 
			:description, 
			:status, 
			:price,
			:image_file,
			:categories,
			:isfavorite
		)`

	_, err := r.NamedExec(queryproduct, data)
	if err != nil {
		return nil, err
	}
	return &config.Result{Message: "1 product user created"}, nil
}

func (r RepoProduct) DeleteProduct(data *model.Product) (*config.Result, error) {
	queryproduct := `DELETE FROM public.products WHERE id_product=:id_product`
	_, err := r.NamedExec(queryproduct, data)
	if err != nil {
		return nil, err
	}
	return &config.Result{Message: "1 product product deleted"}, nil

}

func (r RepoProduct) UpdateProduct(data *model.Product) (*config.Result, error) {
	queryProduct := `UPDATE public.products SET
	product_name = :product_name,
	description = :description,
	status = :status,
	price = :price,
	image_file = :image_file,
	categories = :categories,
	isfavourite = :isfavourite
	WHERE id_product = :id_product
	`

	_, err := r.NamedExec(queryProduct, data)
	fmt.Println(r.NamedExec(queryProduct, data))
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	return &config.Result{Message: "1 product user updated"}, nil
}

func (r RepoProduct) Get_Products(data *model.Product, search string, limit int, page int, orderby string) ([]model.Product, *Pagination, error) {
	products_data := []model.Product{}

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

	sqltable := fmt.Sprintln("SELECT count(id_product) FROM products")

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
		r.Select(&products_data, `select * from products WHERE categories=$1 LIMIT $2 OFFSET $3 `, search, limit, offset)
	} else if orderby != "" {
		r.Select(&products_data, `select * from products order by `+orderby+` asc LIMIT $1 OFFSET $2`, limit, offset)
	} else if orderby != "" && search != "" {
		r.Select(&products_data, `select * from products order by `+orderby+` asc WHERE categories=$1 LIMIT $2 OFFSET $3 `, search, limit, offset)
	} else {
		r.Select(&products_data, `select * from products LIMIT $1 OFFSET $2 `, limit, offset)
	}
	if len(products_data) == 0 {
		return nil, nil, errors.New("data not found")
	}
	return products_data, pgnt, nil
}
