CREATE TABLE public.product (
	id serial4 NOT NULL,
	product_name varchar(50) NOT NULL,
	price numeric(10, 2) NOT NULL,
	CONSTRAINT product_pkey PRIMARY KEY (id)
);