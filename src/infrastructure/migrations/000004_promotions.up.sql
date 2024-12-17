CREATE TABLE public.promotions (
	id_promotion uuid NOT NULL DEFAULT uuid_generate_v4(),
	code varchar(255) NOT NULL UNIQUE,
    description text NOT NULL,
    discount NUMERIC(5, 2) NOT NULL CHECK (discount > 0 AND discount <= 100), 
    -- Giới hạn discount: tối đa 100% (nếu là tỷ lệ), hoặc phù hợp với đơn vị tiền
    quantity int NOT NULL DEFAULT 0,
    created_at timestamp NULL DEFAULT now(),
	updated_at timestamp NULL,
    started_at timestamp NULL DEFAULT now(),
	end_at TIMESTAMP NOT NULL CHECK (end_at > started_at),

	CONSTRAINT promotion_pk PRIMARY KEY (id_promotion)
);