-- Tạo bảng `order_details`
CREATE TABLE public.order_details (
    id_order_detail SERIAL PRIMARY KEY,                   -- ID dòng chi tiết đơn hàng
    id_order uuid NOT NULL REFERENCES public.orders(id_order) ON DELETE CASCADE, -- Liên kết đơn hàng
    item_type varchar(10) NOT NULL CHECK (item_type IN ('product', 'book')), -- Loại mục: sản phẩm hoặc sách
    id_item uuid NOT NULL,                                -- ID của sản phẩm hoặc sách
    quantity int NOT NULL DEFAULT 1,                      -- Số lượng sản phẩm hoặc sách
    price NUMERIC(10, 2),                                 -- Giá của sản phẩm (NULL cho sách mượn)
    note text NULL,                                       -- Ghi chú (ví dụ: "sách mượn tại quán"),

    -- Không dùng CHECK, dùng 2 FOREIGN KEY ràng buộc
    FOREIGN KEY (id_item) REFERENCES public.products(id_product) ON DELETE CASCADE DEFERRABLE INITIALLY DEFERRED,
    FOREIGN KEY (id_item) REFERENCES public.books(id_book) ON DELETE CASCADE DEFERRABLE INITIALLY DEFERRED
);

-- Tạo hàm trigger để kiểm tra item_type và id_item
CREATE OR REPLACE FUNCTION check_item_type()
RETURNS TRIGGER AS $$
BEGIN
    IF (NEW.item_type = 'product' AND NOT EXISTS (SELECT 1 FROM public.products WHERE id_product = NEW.id_item)) THEN
        RAISE EXCEPTION 'Invalid product ID for item_type product';
    ELSIF (NEW.item_type = 'book' AND NOT EXISTS (SELECT 1 FROM public.books WHERE id_book = NEW.id_item)) THEN
        RAISE EXCEPTION 'Invalid book ID for item_type book';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Gắn trigger vào bảng order_details
CREATE TRIGGER trg_check_item_type
BEFORE INSERT OR UPDATE ON public.order_details
FOR EACH ROW
EXECUTE FUNCTION check_item_type();