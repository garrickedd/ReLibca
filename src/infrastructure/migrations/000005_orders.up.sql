CREATE TABLE public.orders (
    id_order uuid NOT NULL DEFAULT uuid_generate_v4(),      -- ID đơn hàng
    id_user uuid NOT NULL REFERENCES public.users(id_user) ON DELETE CASCADE, -- Người tạo đơn
    id_promotion uuid NULL REFERENCES public.promotions(id_promotion), -- Mã giảm giá (nếu có)
    total_amount NUMERIC(10, 2) NOT NULL,                  -- Tổng số tiền thanh toán
    payment_method varchar(50) NOT NULL,                   -- Hình thức thanh toán (e.g., "cash", "card")
    created_at timestamp NULL DEFAULT now(),               -- Thời gian tạo đơn hàng
    updated_at timestamp NULL,
    CONSTRAINT order_pk PRIMARY KEY (id_order)
);