CREATE TABLE "users" (
  "user_id" bigserial PRIMARY KEY,
  "username" varchar(50) NOT NULL,
  "email" varchar(50) NOT NULL,
  "password" varchar(50) NOT NULL,
  "address" varchar(100) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "products" (
  "product_id" bigserial PRIMARY KEY,
  "name" varchar(100) NOT NULL,
  "description" varchar(50) NOT NULL,
  "price" decimal(10,2) NOT NULL,
  "stock" int NOT NULL,
  "category_id" bigserial NOT NULL
);

CREATE TABLE "categories" (
  "category_id" bigserial PRIMARY KEY,
  "name" varchar(50) NOT NULL,
  "description" text
);

CREATE TABLE "orders" (
  "order_id" bigserial PRIMARY KEY,
  "user_id" bigserial NOT NULL,
  "order_date" timestamptz NOT NULL DEFAULT (now()),
  "status" varchar(20) NOT NULL
);

CREATE INDEX "idx_users_email" ON "users" ("email");

CREATE INDEX "idx_users_username" ON "users" ("username");

CREATE INDEX "idx_products_category_id" ON "products" ("category_id");

CREATE INDEX "idx_products_price" ON "products" ("price");

CREATE INDEX "idx_orders_user_id" ON "orders" ("user_id");

CREATE INDEX "idx_orders_order_date" ON "orders" ("order_date");

CREATE INDEX "idx_orders_status" ON "orders" ("status");

ALTER TABLE "products" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("category_id");

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");
