CREATE SEQUENCE main_meal_seq START WITH 1 INCREMENT BY 1;
CREATE TABLE IF NOT EXISTS main_meal
(
    meal_id int DEFAULT nextval
(
    'main_meal_seq'
) PRIMARY KEY, meal_name text );

CREATE SEQUENCE stew_seq START WITH 1 INCREMENT BY 1;
CREATE TABLE IF NOT EXISTS stew
(
    stew_id int DEFAULT nextval
(
    'stew_seq'
) PRIMARY KEY, stew text );

CREATE SEQUENCE image_seq START WITH 1 INCREMENT BY 1;
CREATE TABLE IF NOT EXISTS food_image
(
    image_id int DEFAULT nextval
(
    'image_seq'
) PRIMARY KEY, food_name text, food_image text );

CREATE SEQUENCE exception_seq START WITH 1 INCREMENT BY 1;
CREATE TABLE IF NOT EXISTS stew
(
    exception_id int DEFAULT nextval
(
    'exception_seq'
) PRIMARY KEY, meal_id INT FOREIGN KEY main_meal
(
    meal_id
), stew_id INT FOREIGN KEY stew
(
    stew_id
));
