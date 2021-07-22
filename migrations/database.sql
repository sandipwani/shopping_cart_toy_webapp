DROP DATABASE pgdata;
CREATE DATABASE pgdata WITH OWNER = postgres;

\c pgdata

DROP TABLE product; 

CREATE TABLE product (
	prod_id integer PRIMARY KEY,
	prod_name varchar(100) NOT NULL,
	prod_price float NOT NULL
);

INSERT INTO product (prod_id, prod_name, prod_price) VALUES (101, 'Redmi note 1', 11999.00);
INSERT INTO product (prod_id, prod_name, prod_price) VALUES (102, 'Redmi note 2', 12999.50);
INSERT INTO product (prod_id, prod_name, prod_price) VALUES (103, 'Redmi note 3', 13999.59);
INSERT INTO product (prod_id, prod_name, prod_price) VALUES (104, 'Redmi note 4', 14999.45);
INSERT INTO product (prod_id, prod_name, prod_price) VALUES (105, 'Redmi note 5', 15999.33);
INSERT INTO product (prod_id, prod_name, prod_price) VALUES (106, 'Redmi note 6', 16999);
INSERT INTO product (prod_id, prod_name, prod_price) VALUES (107, 'Redmi note 7', 17999.55);
INSERT INTO product (prod_id, prod_name, prod_price) VALUES (108, 'Redmi note 8', 18999.99);
INSERT INTO product (prod_id, prod_name, prod_price) VALUES (109, 'Redmi note 9', 19999.99);
INSERT INTO product (prod_id, prod_name, prod_price) VALUES (110, 'Redmi note 10', 20999.98);

DROP TABLE product_cart; 

CREATE TABLE product_cart(
    productid integer PRIMARY KEY,
    productname varchar (100) NOT NULL,
    productprice float NOT NULL,
    productquantity integer NOT NULL
);

INSERT INTO product_cart (productid, productname, productprice, productquantity) VALUES (1, 'Microwave Oven', 25999.79, 6);
          


--psql -f /home/webonise/go/src/shopping_cart_app/database/database.sql