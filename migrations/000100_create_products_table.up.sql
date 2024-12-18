CREATE TABLE IF NOT EXISTS "products"
(
    "sku"                   varchar(15),
    "name"                  varchar(200),
    "category"              varchar(50),
    "price"                 int NOT NULL
);

CREATE INDEX IF NOT EXISTS category_price_idx ON products (category, price);

-- test data

INSERT INTO products ("sku", "name", "category", "price") VALUES ('000001', 'BV Lean leather ankle boots', 'boots', 89000);
INSERT INTO products ("sku", "name", "category", "price") VALUES ('000002', 'BV Lean leather ankle boots', 'boots', 99000);
INSERT INTO products ("sku", "name", "category", "price") VALUES ('000003', 'Ashlington leather ankle boots', 'boots', 71000);
INSERT INTO products ("sku", "name", "category", "price") VALUES ('000004', 'Naima embellished suede sandals', 'sandals', 79500);
INSERT INTO products ("sku", "name", "category", "price") VALUES ('000005', 'Nathane leather sneakers', 'sneakers', 59000);