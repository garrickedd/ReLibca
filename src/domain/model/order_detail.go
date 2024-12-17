package model

// id_order_detail SERIAL PRIMARY KEY,                   -- ID dòng chi tiết đơn hàng
// id_order uuid NOT NULL REFERENCES public.orders(id_order) ON DELETE CASCADE, -- Liên kết đơn hàng
// item_type varchar(10) NOT NULL CHECK (item_type IN ('product', 'book')), -- Loại mục: sản phẩm hoặc sách
// id_item uuid NOT NULL,                                -- ID của sản phẩm hoặc sách
// quantity int NOT NULL DEFAULT 1,                      -- Số lượng sản phẩm hoặc sách
// price NUMERIC(10, 2),                                 -- Giá của sản phẩm (NULL cho sách mượn)
// note text NULL,                                       -- Ghi chú (ví dụ: "sách mượn tại quán"),

type OrderDetail struct {
	Id_order_detail int      `db:"id_order_detail" json:"id_order_detail"`
	Id_order        string   `db:"id_order" json:"id_order"`
	Item_type       string   `db:"item_type" json:"item_type"`
	Id_item         string   `db:"id_item" json:"id_item"`
	Quantity        int      `db:"quantity" json:"quantity"`
	Price           *float64 `db:"price" json:"price,omitempty"`
	Note            *string  `db:"note" json:"note,omitempty"`
}
