package repository

import (
	"errors"
	"fmt"

	"github.com/garrickedd/ReLibca/src/domain/model"
	"github.com/garrickedd/ReLibca/src/shared/config"
	"github.com/jmoiron/sqlx"
)

type RepoOrderIF interface {
	CreateOrder(data *model.Order) (*config.Result, error)
	UpdateOrder(data *model.Order) (*config.Result, error)
	DeleteOrder(data *model.Order) (*config.Result, error)
	Get_Orders(data *model.Order, search string, limit int, page int, orderby string) ([]model.Order, *Pagination, error)
}

type RepoOrder struct {
	*sqlx.DB
}

func NewOrder(db *sqlx.DB) *RepoOrder {
	return &RepoOrder{db}
}

// CreateOrder creates a new order and its associated order details
func (r RepoOrder) CreateOrder(data *model.Order) (*config.Result, error) {
	tx := r.DB.MustBegin()

	queryOrder := `INSERT INTO orders(
		id_user,  
		total_amount, 
		discount_amount, 
		payment_method,         
		status,
		created_at,
		updated_at
	) 
		VALUES(
			:id_user, 
			:total_amount, 
			:discount_amount, 
			:payment_method, 
			:status,
			:created_at,
			:updated_at
		) RETURNING id_order`

	err := tx.Get(&data.Id_order, queryOrder, data)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, item := range data.OrderDetails {
		item.Id_order = data.Id_order
		queryOrderDetails := `INSERT INTO order_details(
			id_order, 
			item_type, 
			id_item, 
			quantity, 
			price
		) VALUES(
			:id_order, 
			:item_type, 
			:id_item, 
			:quantity, 
			:price
		)`

		_, err := tx.NamedExec(queryOrderDetails, item)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &config.Result{Message: "1 order created"}, nil
}

// UpdateOrder updates an existing order and its associated order details
func (r RepoOrder) UpdateOrder(data *model.Order) (*config.Result, error) {
	tx := r.DB.MustBegin()

	queryOrder := `UPDATE orders SET
		id_user = :id_user,
		total_amount = :total_amount,
		discount_amount = :discount_amount,
		payment_method = :payment_method,
		status = :status,
		updated_at = :updated_at
	WHERE id_order = :id_order`

	_, err := tx.NamedExec(queryOrder, data)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, item := range data.OrderDetails {
		queryOrderDetails := `UPDATE order_details SET
			item_type = :item_type,
			id_item = :id_item,
			quantity = :quantity,
			price = :price
		WHERE id_order_detail = :id_order_detail`

		_, err := tx.NamedExec(queryOrderDetails, item)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &config.Result{Message: "1 order updated"}, nil
}

// DeleteOrder deletes an order and its associated order details
func (r RepoOrder) DeleteOrder(data *model.Order) (*config.Result, error) {
	queryOrder := `DELETE FROM orders WHERE id_order=:id_order`
	queryOrderDetails := `DELETE FROM order_details WHERE id_order=:id_order`

	tx := r.DB.MustBegin()

	_, err := tx.Exec(queryOrderDetails, data)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	_, err = tx.Exec(queryOrder, data)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &config.Result{Message: "1 order deleted"}, nil
}

// Get_Orders retrieves orders based on search criteria, pagination, and ordering
func (r RepoOrder) Get_Orders(data *model.Order, search string, limit int, page int, orderby string) ([]model.Order, *Pagination, error) {
	orders_data := []model.Order{}

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

	sqltable := fmt.Sprintln("SELECT count(id_order) FROM orders")

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
		r.Select(&orders_data, `select * from orders WHERE id_order=$1 LIMIT $2 OFFSET $3`, search, limit, offset)
	} else if orderby != "" {
		r.Select(&orders_data, `select * from orders order by `+orderby+` asc LIMIT $1 OFFSET $2`, limit, offset)
	} else if orderby != "" && search != "" {
		r.Select(&orders_data, `select * from orders order by `+orderby+` asc WHERE id_order=$1 LIMIT $2 OFFSET $3`, search, limit, offset)
	} else {
		r.Select(&orders_data, `select * from orders LIMIT $1 OFFSET $2`, limit, offset)
	}
	if len(orders_data) == 0 {
		return nil, nil, errors.New("data not found")
	}
	return orders_data, pgnt, nil
}
