CREATE TABLE wallet (
    id serial primary key,
    user_id int unique,
    name varchar(50),
    balance int,
    Foreign Key (user_id) REFERENCES users(id),
    update_at timestamp default current_timestamp,
    created_at timestamp default current_timestamp
)