CREATE TABLE orders (
    order_uid VARCHAR(30) PRIMARY KEY,
    track_number VARCHAR(30) NOT NULL,
    entry VARCHAR(30),
    locale VARCHAR(10),
    internal_signature VARCHAR(30),
    customer_id VARCHAR(30) NOT NULL,
    delivery_service VARCHAR(30) NOT NULL,
    shardkey VARCHAR(30),
    sm_id int NOT NULL,
    date_created VARCHAR(40) NOT NULL,
    oof_shard VARCHAR(30)
);

CREATE TABLE payments (               
    transaction VARCHAR(40) PRIMARY KEY,
    request_id VARCHAR(30),
    currency VARCHAR(30) NOT NULL,
    provider VARCHAR(30) NOT NULL,
    amount int NOT NULL,
    payment_dt int NOT NULL,
    bank VARCHAR(30) NOT NULL,
    delivery_cost int NOT NULL,
    goods_total int NOT NULL,
    custom_fee int,
    fk_payments_order VARCHAR(30) REFERENCES orders(order_uid) ON DELETE CASCADE
);

CREATE TABLE delivery (
    delivery_id serial PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    phone VARCHAR(30) NOT NULL,
    zip VARCHAR(30) NOT NULL,
    city VARCHAR(30) NOT NULL,
    address VARCHAR(40) NOT NULL,
    region VARCHAR(30),
    email VARCHAR(30) NOT NULL,
    fk_delivery_order VARCHAR(30) REFERENCES orders(order_uid) ON DELETE CASCADE
);

CREATE TABLE items (
    chrt_id int PRIMARY KEY,
    track_number VARCHAR(30) NOT NULL,
    price int NOT NULL,
    rid VARCHAR(30) NOT NULL,
    name VARCHAR(30) NOT NULL,
    sale int,
    size VARCHAR,
    total_price int NOT NULL,
    nm_id int NOT NULL,
    brand VARCHAR(30),
    status int NOT NULL,
    fk_items_order VARCHAR(30) REFERENCES orders(order_uid) ON DELETE CASCADE
);