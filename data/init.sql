CREATE TABLE products
(
    id SERIAL,
    name TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    CONSTRAINT products_pkey PRIMARY KEY (id)
);

INSERT INTO products (name, price) VALUES ('APPLE iPhone 12', 899.99);
INSERT INTO products (name, price) VALUES ('APPLE iPhone 12 Pro', 1089);
INSERT INTO products (name, price) VALUES ('APPLE iPhone 11', 629.99);
INSERT INTO products (name, price) VALUES ('SAMSUNG Galaxy A52', 345.99);
INSERT INTO products (name, price) VALUES ('SAMSUNG Galaxy S20 FE', 625);
INSERT INTO products (name, price) VALUES ('GOOGLE Pixel 4a', 333);
INSERT INTO products (name, price) VALUES ('XIAOMI Mi Not 10 lite', 399);