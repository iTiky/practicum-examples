create extension if not exists "uuid-ossp";

-- Users table
create table users
(
    id         uuid                  default uuid_generate_v4(),
    name       varchar(100) not null,
    email      varchar(100) not null,
    phone      varchar(100),
    region     varchar(5)   not null,
    created_at timestamp    not null default now(),
    updated_at timestamp    not null default now(),
    deleted_at timestamp,
    primary key (id),
    unique (email)
);

create index users_email_idx ON users (email);

-- Orders table
create table orders
(
    id         uuid               default uuid_generate_v4(),
    user_id    uuid      not null,
    status     int       not null,
    created_at timestamp not null default now(),
    deleted_at timestamp,
    primary key (id),
    foreign key (user_id) references users (id)
);

create index orders_user_id_idx on orders (user_id);
