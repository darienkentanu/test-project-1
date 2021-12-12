create database if not exists `KlikA2C`;
use `KlikA2C`;
create table if not exists users(
id int primary key auto_increment,
name varchar(100) not null,
email varchar(100) not null unique,
password varchar(255) not null,
token varchar(255),
created_at DATETIME(3),
updated_at DATETIME(3),
deleted_at DATETIME(3)
);
create table if not exists items(
id int AUTO_INCREMENT PRIMARY KEY, 
name VARCHAR(100) NOT NULL, 
price DECIMAL(20,2) NOT NULL, 
cost DECIMAL(20,2) NOT NULL,
created_at DATETIME(3),
updated_at DATETIME(3)
);
create table if not exists transactions(
id int primary key auto_increment,
number int(8) not null unique,
price_total decimal(20,2) not null,
cost_total decimal(20,2) not null,
created_at DATETIME(3),
updated_at DATETIME(3),
deleted_at DATETIME(3)
);
create table if not exists transaction_details(
id int primary key auto_increment,
transaction_id int,
item_id int,
item_quantity int not null,
item_price decimal(20,2) not null,
item_cost decimal(20,2) not null,
created_at DATETIME(3),
updated_at DATETIME(3),
deleted_at DATETIME(3),
CONSTRAINT fk_transaction_details_transaction 
FOREIGN KEY (transaction_id)
REFERENCES transactions(id),
CONSTRAINT fk_transaction_details_items
FOREIGN KEY (item_id)
REFERENCES items(id)
);
create database if not exists `KlikA2C_Test`;
use `KlikA2C_Test`;
create table if not exists users(
id int primary key auto_increment,
name varchar(100) not null,
email varchar(100) not null unique,
password varchar(255) not null,
token varchar(255),
created_at DATETIME(3),
updated_at DATETIME(3),
deleted_at DATETIME(3)
);
create table if not exists items(
id int AUTO_INCREMENT PRIMARY KEY, 
name VARCHAR(100) NOT NULL, 
price DECIMAL(20,2) NOT NULL, 
cost DECIMAL(20,2) NOT NULL,
created_at DATETIME(3),
updated_at DATETIME(3)
);
create table if not exists transactions(
id int primary key auto_increment,
number int(8) not null unique,
price_total decimal(20,2) not null,
cost_total decimal(20,2) not null,
created_at DATETIME(3),
updated_at DATETIME(3),
deleted_at DATETIME(3)
);
create table if not exists transaction_details(
id int primary key auto_increment,
transaction_id int,
item_id int,
item_quantity int not null,
item_price decimal(20,2) not null,
item_cost decimal(20,2) not null,
created_at DATETIME(3),
updated_at DATETIME(3),
deleted_at DATETIME(3),
CONSTRAINT fk_transaction_details_transaction 
FOREIGN KEY (transaction_id)
REFERENCES transactions(id),
CONSTRAINT fk_transaction_details_items
FOREIGN KEY (item_id)
REFERENCES items(id)
);