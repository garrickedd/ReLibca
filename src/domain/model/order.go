package model

import "time"

// id_order uuid NOT NULL DEFAULT uuid_generate_v4(),      -- ID đơn hàng
// id_user uuid NOT NULL REFERENCES public.users(id_user) ON DELETE CASCADE, -- Người tạo đơn
// id_promotion uuid NULL REFERENCES public.promotions(id_promotion), -- Mã giảm giá (nếu có)
// total_amount NUMERIC(10, 2) NOT NULL,                  -- Tổng số tiền thanh toán
// payment_method varchar(50) NOT NULL,                   -- Hình thức thanh toán (e.g., "cash", "card")
// created_at timestamp NULL DEFAULT now(),               -- Thời gian tạo đơn hàng
// updated_at timestamp NULL,

type Order struct {
	Id_order       string     `db:"id_order" json:"id_order"`
	Id_user        string     `db:"id_user" json:"id_user"`
	Id_promotion   *string    `db:"id_promotion" json:"id_promotion,omitempty"`
	Total_amount   float64    `db:"total_amount" json:"total_amount"`
	Payment_method string     `db:"payment_method" json:"payment_method"`
	Created_at     *time.Time `db:"created_at" json:"created_at"`
	Updated_at     *time.Time `db:"updated_at" json:"updated_at,omitempty"`
}
