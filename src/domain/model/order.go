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
	Id_order        string        `db:"id_order" form:"id_order" json:"id_order"`
	Id_user         string        `db:"id_user" form:"id_user" json:"id_user"`
	Total_amount    float64       `db:"total_amount" form:"total_amount" json:"total_amount"`
	Discount_amount float64       `db:"discount_amount" form:"discount_amount" json:"discount_amount"`
	Payment_method  string        `db:"payment_method" form:"payment_method" json:"payment_method"`
	Status          string        `db:"status" form:"status" json:"status"`
	Created_at      *time.Time    `db:"created_at" form:"created_at" json:"created_at"`
	Updated_at      *time.Time    `db:"updated_at" form:"updated_at" json:"updated_at"`
	OrderDetails    []OrderDetail `db:"order_details" form:"order_details" json:"order_details"`
}

// id_order_detail SERIAL PRIMARY KEY,                   -- ID dòng chi tiết đơn hàng
// id_order uuid NOT NULL REFERENCES public.orders(id_order) ON DELETE CASCADE, -- Liên kết đơn hàng
// item_type varchar(10) NOT NULL CHECK (item_type IN ('product', 'book')), -- Loại mục: sản phẩm hoặc sách
// id_item uuid NOT NULL,                                -- ID của sản phẩm hoặc sách
// quantity int NOT NULL DEFAULT 1,                      -- Số lượng sản phẩm hoặc sách
// price NUMERIC(10, 2),                                 -- Giá của sản phẩm (NULL cho sách mượn)
// note text NULL,                                       -- Ghi chú (ví dụ: "sách mượn tại quán"),
type OrderDetail struct {
	Id_order_detail string  `db:"id_order_detail" form:"id_order_detail" json:"id_order_detail"`
	Id_order        string  `db:"id_order" form:"id_order" json:"id_order"`
	Item_type       string  `db:"item_type" form:"item_type" json:"item_type"`
	Id_item         string  `db:"id_item" form:"id_item" json:"id_item"`
	Quantity        int     `db:"quantity" form:"quantity" json:"quantity"`
	Price           float64 `db:"price" form:"price" json:"price"`
}
