package model

import "time"

// id_promotion uuid NOT NULL DEFAULT uuid_generate_v4(),
// code varchar(255) NOT NULL UNIQUE,
// description text NOT NULL,
// discount NUMERIC(5, 2) NOT NULL CHECK (discount > 0 AND discount <= 100),
// -- Giới hạn discount: tối đa 100% (nếu là tỷ lệ), hoặc phù hợp với đơn vị tiền
// quantity int NOT NULL DEFAULT 0,
// created_at timestamp NULL DEFAULT now(),
// updated_at timestamp NULL,
// started_at timestamp NULL DEFAULT now(),
// end_at TIMESTAMP NOT NULL CHECK (end_at > started_at),

type Promotion struct {
	Id_promotion string     `db:"id_promotion" form:"id_promotion" json:"id_promotion"`
	Code         string     `db:"code" form:"code" json:"code"`
	Description  string     `db:"description" form:"description" json:"description"`
	Discount     float64    `db:"discount" form:"discount" json:"discount"`
	Quantity     int        `db:"quantity" form:"quantity" json:"quantity"`
	Created_at   *time.Time `db:"created_at" json:"created_at"`
	Updated_at   *time.Time `db:"updated_at" json:"updated_at,omitempty"`
	Started_at   *time.Time `db:"started_at" json:"started_at"`
	End_at       *time.Time `db:"end_at" json:"end_at"`
}
