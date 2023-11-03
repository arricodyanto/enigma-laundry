CREATE TABLE mst_customer (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    contact VARCHAR(15) NOT NULL
);

CREATE TABLE mst_service (
    id SERIAL PRIMARY KEY,
    service VARCHAR(100) NOT NULL,
    unit VARCHAR(7) NOT NULL,
    price INT NOT NULL
);

CREATE TABLE trx_bill (
    id SERIAL PRIMARY KEY,
    customer_id INT NOT NULL,
    entry_date DATE NOT NULL,
    out_date DATE NOT NULL,
    recipient_name VARCHAR(50) NOT NULL,
    total_bill INT,
    FOREIGN KEY (customer_id) REFERENCES mst_customer (id) ON DELETE CASCADE
);

CREATE TABLE trx_bill_detail (
    id SERIAL PRIMARY KEY,
    bill_id INT NOT NULL,
    service_id INT NOT NULL,
    amount INT NOT NULL,
    total INT NOT NULL,
    FOREIGN KEY (bill_id) REFERENCES trx_bill (id) ON DELETE CASCADE,
    FOREIGN KEY (service_id) REFERENCES mst_service (id) ON DELETE CASCADE
);
