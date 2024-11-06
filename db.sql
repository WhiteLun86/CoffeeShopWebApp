create table coffee_products
(
    id          int auto_increment
        primary key,
    name        varchar(255)   not null,
    price       decimal(10, 2) not null,
    description text           null,
    Stock       int default 0  not null
);

create table users
(
    id       int auto_increment
        primary key,
    username varchar(50)            not null,
    password varchar(100)           not null,
    role     enum ('user', 'admin') not null,
    constraint username
        unique (username)
);

create table orders
(
    id          int auto_increment
        primary key,
    user_id     int                                                         not null,
    total_price decimal(10, 2)                                              not null,
    status      enum ('pending', 'completed', 'canceled') default 'pending' null,
    constraint orders_ibfk_1
        foreign key (user_id) references users (id)
            on delete cascade
);

create table order_products
(
    order_id   int            not null,
    product_id int            not null,
    quantity   int            not null,
    price      decimal(10, 2) not null,
    primary key (order_id, product_id),
    constraint order_products_ibfk_1
        foreign key (order_id) references orders (id)
            on delete cascade,
    constraint order_products_ibfk_2
        foreign key (product_id) references coffee_products (id)
            on delete cascade
);

create index product_id
    on order_products (product_id);

