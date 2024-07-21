CREATE TABLE users(
    id serial primary key,
    username varchar(50) not null,
    email varchar(50) not null,
    password varchar(50) not null,
    create_at timestamp default current_timestamp
)