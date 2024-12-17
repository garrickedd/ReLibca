package repository

import (
	"errors"
	"fmt"
	"log"

	"github.com/garrickedd/ReLibca/src/domain/model"
	"github.com/garrickedd/ReLibca/src/shared/config"
	"github.com/jmoiron/sqlx"
)

type RepoPromotionIF interface {
	CreatePromotion(data *model.Promotion) (*config.Result, error)
	UpdatePromotion(data *model.Promotion) (*config.Result, error)
	DeletePromotion(data *model.Promotion) (*config.Result, error)
	Get_Promotions(data *model.Promotion, search string, limit int, page int, orderby string) ([]model.Promotion, *Pagination, error)
}

type RepoPromotion struct {
	*sqlx.DB
}

func NewPromotion(db *sqlx.DB) *RepoPromotion {
	return &RepoPromotion{db}
}

func (r RepoPromotion) CreatePromotion(data *model.Promotion) (*config.Result, error) {
	queryPromotion := `INSERT INTO promotions(
		code,  
		description, 
		discount,         
		quantity,
		started_at,
		end_at
	) 
		VALUES(
			:code, 
			:description, 
			:discount, 
			:quantity,
			:started_at,
			:end_at
		)`

	_, err := r.NamedExec(queryPromotion, data)
	if err != nil {
		return nil, err
	}
	return &config.Result{Message: "1 promotion created"}, nil
}

func (r RepoPromotion) DeletePromotion(data *model.Promotion) (*config.Result, error) {
	queryPromotion := `DELETE FROM public.promotions WHERE id_promotion=:id_promotion`
	_, err := r.NamedExec(queryPromotion, data)
	if err != nil {
		return nil, err
	}
	return &config.Result{Message: "1 promotion deleted"}, nil
}

func (r RepoPromotion) UpdatePromotion(data *model.Promotion) (*config.Result, error) {
	queryPromotion := `UPDATE public.promotions SET
	code = :code,
	description = :description,
	discount = :discount,
	quantity = :quantity,
	started_at = :started_at,
	end_at = :end_at
	WHERE id_promotion = :id_promotion
	`

	_, err := r.NamedExec(queryPromotion, data)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	return &config.Result{Message: "1 promotion updated"}, nil
}

func (r RepoPromotion) Get_Promotions(data *model.Promotion, search string, limit int, page int, orderby string) ([]model.Promotion, *Pagination, error) {
	promotions_data := []model.Promotion{}

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

	sqltable := fmt.Sprintln("SELECT count(id_promotion) FROM promotions")

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
		r.Select(&promotions_data, `select * from promotions WHERE code=$1 LIMIT $2 OFFSET $3 `, search, limit, offset)
	} else if orderby != "" {
		r.Select(&promotions_data, `select * from promotions order by `+orderby+` asc LIMIT $1 OFFSET $2`, limit, offset)
	} else if orderby != "" && search != "" {
		r.Select(&promotions_data, `select * from promotions order by `+orderby+` asc WHERE code=$1 LIMIT $2 OFFSET $3 `, search, limit, offset)
	} else {
		r.Select(&promotions_data, `select * from promotions LIMIT $1 OFFSET $2 `, limit, offset)
	}
	if len(promotions_data) == 0 {
		return nil, nil, errors.New("data not found")
	}
	return promotions_data, pgnt, nil
}
