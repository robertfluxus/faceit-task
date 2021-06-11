CREATE SCHEMA IF NOT EXISTS faceit;

CREATE TABLE faceit.users (
    id varchar NOT NULL PRIMARY KEY,
    request_id varchar NOT NULL UNIQUE,
    first_name varchar NOT NULL,
    last_name varchar NOT NULL,
    nickname varchar NOT NULL UNIQUE,
    password varchar NOT NULL,
    email varchar NOT NULL,
    country varchar NOT NULL
);

CREATE INDEX user_country_idx ON faceit.users (country);
CREATE INDEX user_nickname_idx ON faceit.users (nickname);

INSERT INTO faceit.users (
    id,
    request_id,
    first_name,
    last_name,
    nickname,
    password,
    email,
    country
) VALUES (
    'testuser',
    'initialrequest',
    'Robert',
    'Marincu',
    'rob_',
    'mysecretpassword',
    'rsmarincu@gmail.com',
    'Romania'
);