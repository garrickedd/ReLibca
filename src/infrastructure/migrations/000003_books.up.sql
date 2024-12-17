CREATE TABLE public.books (
    id_book uuid NOT NULL DEFAULT uuid_generate_v4(),
	title varchar(255) NOT NULL,
    author varchar(255) NOT NULL,
    count int NOT NULL DEFAULT 0,
    tag text[] NOT NULL,
    place varchar(255),
    image_file varchar NULL,
	CONSTRAINT book_pk PRIMARY KEY (id_book)
);