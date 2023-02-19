DROP TABLE IF EXISTS create table product(
    product_id INT GENERATED ALWAYS AS IDENTITY,
    product_name varchar(100) Unique,
    date_created timestamp default CURRENT_TIMESTAMP,
    serial_number varchar(100),
    product_quantity integer,
    product_price integer,
    product_image text,
    category_id integer,

    PRIMARY KEY(product_id),

     CONSTRAINT fk_category
      FOREIGN KEY(category_id) 
	  REFERENCES category(category_id)
	  ON DELETE CASCADE
);


 create table category(
    category_id INT GENERATED ALWAYS AS IDENTITY,
    name varchar(100),
    date_created timestamp default current_timestamp,
    PRIMARY KEY(category_id)
);
create table transaction_logs(
    transaction_id INT GENERATED ALWAYS AS IDENTITY,
    product_name varchar(100),
    transaction_time timestamp default CURRENT_TIMESTAMP,
    product_price integer,

    primary key (transaction_id),

    CONSTRAINT fk_product
      FOREIGN KEY(product_name) 
	  REFERENCES product(product_name)

    CONSTRAINT fk_product
      FOREIGN KEY(product_price) 
	  REFERENCES product(product_price)

);