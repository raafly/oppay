-- Active: 1723170566694@@127.0.0.1@5432@ewallet
CREATE TABLE users (
    id varchar(36) primary key,
    name varchar(50),
    email varchar(50) unique not null,
    pin int not null,
    verification_token varchar(60)
    email_verified_at timestamp,
    remember_token varchar(36),
    updated_at timestamp,
    create_at timestamp default current_timestamp
)

DROP TABLE users