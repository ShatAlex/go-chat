CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE chats
(
    id serial not null unique,
    name varchar(255) not null unique
);

CREATE TABLE chats_of_users
(
    id serial not null unique,
    chat_id int references chats(id) on delete cascade not null,
    user_id int references users(id) on delete cascade not null
);

CREATE TABLE messages
(
    id serial not null unique,
    user_id int references users(id) on delete cascade not null,
    chat_id int references chats(id) on delete cascade not null,
    content varchar(255)
);