CREATE TABLE users (
    id serial PRIMARY KEY,
    name VARCHAR ( 50 ) NOT NULL,
    surname VARCHAR ( 50 ) NOT NULL,
    password VARCHAR ( 50 ) NOT NULL,
    email VARCHAR ( 255 ) UNIQUE NOT NULL
);