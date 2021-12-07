CREATE TABLE merchants(
	id bigserial PRIMARY KEY,
	name varchar(191) NOT NULL,
	lat NUMERIC(10, 6) NOT NULL,
	long NUMERIC(10, 6) NOT NULL,
	logo varchar(191) NOT NULL,
	created_at timestamptz NOT NULL DEFAULT (now()),
	updated_at timestamptz NOT NULL DEFAULT (now())
);


CREATE TABLE categories(
	id bigserial PRIMARY KEY,
	name varchar(191) NOT NULL,
	created_at timestamptz NOT NULL DEFAULT (now()),
	updated_at timestamptz NOT NULL DEFAULT (now())
);


CREATE TABLE products(
	id bigserial PRIMARY KEY,
	name varchar(191) NOT NULL,
	price int NOT NULL,
	category_id int NOT NULL,
	merchant_id int NOT NULL,
	image varchar(191) NOT NULL,
	stock int NOT NULL,
	created_at timestamptz NOT NULL DEFAULT (now()),
	updated_at timestamptz NOT NULL DEFAULT (now()),
		CONSTRAINT products_categories_fk
      		FOREIGN KEY(category_id) 
	      		REFERENCES categories(id)
	        		ON DELETE SET NULL,
		CONSTRAINT products_merchants_fk
      		FOREIGN KEY(merchant_id) 
	      		REFERENCES merchants(id)
	        		ON DELETE SET NULL
);

CREATE TABLE users(
	id bigserial PRIMARY KEY,
	name varchar(191) NOT NULL,
	password varchar(255) NOT NULL,
	email varchar(191) UNIQUE NOT NULL,
	created_at timestamptz NOT NULL DEFAULT (now()),
	updated_at timestamptz NOT NULL DEFAULT (now())
);


CREATE TABLE carts (
	id bigserial PRIMARY KEY,
	user_id int NOT NULL,
	merchant_id int NOT NULL,
	product_id int NOT NULL,
	qty int NOT NULL,
	total_price int NOT NULL,
	created_at timestamptz NOT NULL DEFAULT (now()),
	updated_at timestamptz NOT NULL DEFAULT (now()),
		CONSTRAINT carts_users_fk
      		FOREIGN KEY(user_id) 
	      		REFERENCES users(id)
	        		ON DELETE SET NULL,
		CONSTRAINT carts_merchants_fk
      		FOREIGN KEY(merchant_id) 
	      		REFERENCES merchants(id)
	        		ON DELETE SET NULL
);