INSERT INTO mst_customer (id, name, contact) VALUES
(1, 'Jessica', '0812654987');

INSERT INTO mst_service (id, service, unit, price) VALUES
(1, 'Cuci + Setrika', 'KG', 7000),
(2, 'Laundry Bedcover', 'Buah', 50000),
(3, 'Laundry Boneka', 'Buah', 25000);

INSERT INTO trx_bill (id, customer_id, entry_date, out_date, recipient_name, total_bill) VALUES
(1, 1, '2022-08-18', '2022-08-20', 'Mirna', 135000);

INSERT INTO trx_bill_detail (id, bill_id, service_id, amount, total) VALUES
(1, 1, 1, 5, 35000),
(2, 1, 2, 1, 50000),
(3, 1, 3, 2, 50000);