CREATE TABLE person (  
    userid serial PRIMARY KEY,
    id INT NOT NULL,
    firstname VARCHAR ( 255 ) NOT NULL,
    lastname VARCHAR ( 255 ) NOT NULL,
    address VARCHAR ( 255 ) NOT NULL
);