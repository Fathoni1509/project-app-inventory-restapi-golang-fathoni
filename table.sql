CREATE TABLE users (
	user_id SERIAL PRIMARY KEY NOT NULL,
	username VARCHAR(100) NOT NULL,
	email VARCHAR(100) NOT NULL,
	password VARCHAR(16) NOT NULL,
	role VARCHAR(20) NOT NULL,
	created_at TIMESTAMPTZ
)

CREATE TABLE warehouses (
	warehouse_id SERIAL PRIMARY KEY NOT NULL,
	name VARCHAR(100) NOT NULL,
	location VARCHAR(100) NOT NULL
)

CREATE TABLE shelves (
	shelve_id SERIAL PRIMARY KEY NOT NULL,
	warehouse_id INT NOT NULL,
	CONSTRAINT fk_warehouse
		FOREIGN KEY (warehouse_id)
		REFERENCES warehouses(warehouse_id),
	name VARCHAR(100) NOT NULL
)

CREATE TABLE categories (
	category_id SERIAL PRIMARY KEY NOT NULL,
	name VARCHAR(100) NOT NULL,
	description VARCHAR(256)
)

CREATE TABLE products (
	product_id SERIAL PRIMARY KEY NOT NULL,
	name VARCHAR(100) NOT NULL,
	category_id INT NOT NULL,
	CONSTRAINT fk_category
		FOREIGN KEY (category_id)
		REFERENCES categories(category_id),
	purchase_price NUMERIC,
	sell_price NUMERIC,
	last_updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_by INT NOT NULL,
	CONSTRAINT fk_updated_by
		FOREIGN KEY (updated_by)
		REFERENCES users(user_id)
)

CREATE TABLE inventory (
	inventory_id SERIAL PRIMARY KEY NOT NULL,
	product_id INT NOT NULL,
	CONSTRAINT fk_product
		FOREIGN KEY (product_id)
		REFERENCES products(product_id),
	shelve_id INT NOT NULL,
	CONSTRAINT fk_shelve
		FOREIGN KEY (shelve_id)
		REFERENCES shelves(shelve_id),
	last_updated TIMESTAMPTZ NOT NULL DEFAULT NOW()
)

CREATE TABLE sales (
	sale_id SERIAL PRIMARY KEY NOT NULL,
	user_id INT NOT NULL,
	CONSTRAINT fk_user
		FOREIGN KEY (user_id)
		REFERENCES users(user_id),
	total_purchase NUMERIC,
	sale_date TIMESTAMP NOT NULL
)

CREATE TABLE sales (
	sale_id SERIAL PRIMARY KEY NOT NULL,
	user_id INT NOT NULL,
	CONSTRAINT fk_user
		FOREIGN KEY (user_id)
		REFERENCES users(user_id),
	product_id INT NOT NULL,
	CONSTRAINT fk_produt
		FOREIGN KEY (product_id)
		REFERENCES products(product_id),
	items INT NOT NULL,
	price NUMERIC,
	total NUMERIC,
	sale_date TIMESTAMP NOT NULL
)

CREATE TABLE sales_detail (
	sale_detail_id SERIAL PRIMARY KEY NOT NULL,
	sale_id INT NOT NULL,
	CONSTRAINT fk_sale
		FOREIGN KEY (sale_id)
		REFERENCES sales(sale_id),
	product_id INT NOT NULL,
	CONSTRAINT fk_product
		FOREIGN KEY (product_id)
		REFERENCES products(product_id),
	total_items INT NOT NULL,
	price_per_item NUMERIC,
	subtotal NUMERIC
)

CREATE TABLE sessions (
	session_id SERIAL PRIMARY KEY NOT NULL,
	user_id INT NOT NULL,
	CONSTRAINT fk_user
		FOREIGN KEY (user_id)
		REFERENCES users(user_id),
	expired_at TIMESTAMPTZ NOT NULL,
	revoked_at TIMESTAMPTZ
)

-- alter table
-- Memperbaiki tabel inventory (menambah kolom jumlah)
ALTER TABLE inventory ADD COLUMN quantity INT NOT NULL DEFAULT 0;

-- (Opsional) Memperpanjang kolom password jika nanti ingin pakai Hash (Bcrypt biasanya butuh 60 char)
ALTER TABLE users ALTER COLUMN password TYPE VARCHAR(255);