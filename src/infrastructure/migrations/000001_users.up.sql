CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE public.users (
	id_user uuid NOT NULL DEFAULT uuid_generate_v4(),
	username varchar(255) NOT NULL,
    first_name varchar(15) NOT NULL,
    last_name varchar(15) NOT NULL,
	email varchar(255) NOT NULL,
	pass varchar(255) NOT NULL,
	phone varchar(255) NOT NULL,
	role int NOT NULL DEFAULT 0,
	created_at timestamp NULL DEFAULT now(),
	updated_at timestamp NULL,
	CONSTRAINT user_pk PRIMARY KEY (id_user),
	CONSTRAINT users_un UNIQUE (username)
	-- CONSTRAINT users_username_key UNIQUE (username),
	-- CONSTRAINT users_username_key1 UNIQUE (username)
);