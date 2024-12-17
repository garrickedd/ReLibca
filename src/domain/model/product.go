package model

import "time"

// id_product uuid NOT NULL DEFAULT uuid_generate_v4(),
// 	product_name varchar(255) NOT NULL,
//     description varchar NOT NULL,
//     status bool NOT NULL DEFAULT true,
// 	price int4 NOT NULL,
//     image_file varchar NULL,
//     created_at timestamp NULL DEFAULT now(),
// 	updated_at timestamp NULL,
//     categories varchar NULL DEFAULT 'coffee'::character varying,
//     isfavourite bool NOT NULL DEFAULT false,
// 	CONSTRAINT product_pk PRIMARY KEY (id_product)

type Product struct {
	Id_product   string     `db:"id_product" form:"id_product" json:"id_product" `
	Product_name string     `db:"product_name" form:"product_name" json:"product_name" `
	Description  string     `db:"description" form:"description" json:"description" `
	Status       string     `db:"status" form:"status" json:"status"`
	Price        int64      `db:"price" form:"price" json:"price"`
	Image_file   string     `db:"image_file" json:"image_file,omitempty" valid:"-" `
	Categories   string     `db:"categories" form:"categories" json:"categories" `
	IsFavourite  string     `db:"isfavourite" form:"isfavorite" json:"isfavorite" `
	Created_at   *time.Time `db:"created_at" json:"created_at" `
	Updated_at   *time.Time `db:"updated_at" json:"updated_at" `
}
