CREATE TABLE users(
    id_user int not null primary key AUTO_INCREMENT,
    username varchar(40) not null unique,
    email varchar(30) not null unique,
    encrypted_password LONGTEXT not null
);

CREATE TABLE customer(
    id_customer INT not null PRIMARY KEY AUTO_INCREMENT,
    id_user INT not null unique,
    wallet int not null,
    FOREIGN KEY (id_user) REFERENCES users(id_user) ON DELETE CASCADE
);

CREATE TABLE musician(
    id_mus INT not null PRIMARY KEY AUTO_INCREMENT,
    mus_name varchar(100) not null
);

create table genre(
	id_genre INT not null PRIMARY KEY AUTO_INCREMENT,
	name varchar(50) not null
);

create table preferences(
	id_preference int not null PRIMARY KEY AUTO_INCREMENT,
    id_genre INT not null,
    id_customer int not null,
    FOREIGN KEY (id_genre) REFERENCES genre(id_genre) ON DELETE CASCADE,
    FOREIGN KEY (id_customer) REFERENCES customer(id_customer) ON DELETE CASCADE
);

create table customer_status(
	id_status INT not null PRIMARY KEY AUTO_INCREMENT,
	title varchar(50) not null 
);

create table status_of_customer(
	id_status INT not null,
	id_customer int not null,
    start_time date not null,
    end_time date 
);

create table customer_sale(
	id_customer INT not null,
    sale int not null, 
    customer_rew_count int not null, 
    customer_prob_count int not null
);

Alter table customer_sale ADD FOREIGN KEY (id_customer) REFERENCES customer(id_customer) ON DELETE CASCADE;

create table rewiews(
	id_rew INT not null PRIMARY KEY AUTO_INCREMENT,
    text varchar(100) not null, 
    grade int not null,
    id_mus int not null,
    id_rec int not null,
    id_customer int not null,
    FOREIGN KEY (id_customer) REFERENCES customer(id_customer)ON DELETE CASCADE,  
    FOREIGN KEY (id_mus) REFERENCES musician(id_mus) ON DELETE CASCADE
);

create table records(
	id_rec INT not null PRIMARY KEY AUTO_INCREMENT,
    title varchar(100) not null, 
    duration int not null,
    upload datetime not null,
    id_genre int not null, 
    id_mus int not null,
    bought int,
    FOREIGN KEY (id_mus) REFERENCES musician(id_mus) ON DELETE CASCADE,
    FOREIGN KEY (id_genre) REFERENCES genre(id_genre) ON DELETE CASCADE
);

Alter table rewiews ADD FOREIGN KEY (id_rec) REFERENCES records(id_rec) ON DELETE CASCADE;

create table order_status(
    id_status INT not null PRIMARY KEY AUTO_INCREMENT,
    tittle varchar(100) not null
);

create table order_(
	id_order INT not null PRIMARY KEY AUTO_INCREMENT,
    id_customer INT not null,
    FOREIGN KEY (id_customer) REFERENCES customer(id_customer) ON DELETE CASCADE
);

create table order_info(
    id_order int not null,
    id_status int not null,
    start_day date not null,
    FOREIGN KEY (id_status) REFERENCES order_status(id_status) ON DELETE CASCADE,
    FOREIGN KEY (id_order) REFERENCES order_(id_order) ON DELETE CASCADE
);

create table salon(
	id_salon int not null PRIMARY KEY AUTO_INCREMENT,
    title varchar(45) not null
);

create table order_detail_records(
	id_order INT not null,
    count_ int not null,
    id_rec int not null,
    id_salon int not null,
    FOREIGN KEY (id_rec) REFERENCES records(id_rec) ON DELETE CASCADE,
    FOREIGN KEY (id_order) REFERENCES order_(id_order) ON DELETE CASCADE,
    FOREIGN KEY (id_salon) REFERENCES salon(id_salon) ON DELETE CASCADE
);

create table collections_value(
    id_value int not null PRIMARY KEY AUTO_INCREMENT,
    title varchar(45),
    count_max int not null
);

create table collections(
	id_collection int PRIMARY KEY AUTO_INCREMENT,
    title varchar(45),
    id_value int not null,
    FOREIGN KEY (id_value) REFERENCES collections_value(id_value) ON DELETE CASCADE
);

create table order_detail_collections(
	id_order INT not null,
    count_ int not null,
    id_collection int not null,
    id_salon int not null,
    FOREIGN KEY (id_collection) REFERENCES collections(id_collection) ON DELETE CASCADE,
    FOREIGN KEY (id_order) REFERENCES order_(id_order) ON DELETE CASCADE,
    FOREIGN KEY (id_salon) REFERENCES salon(id_salon) ON DELETE CASCADE
);

create table salon_sales(
	sale_procent int not null,
	date_begin date,
    date_end date,
	id_salon int,
    id_genre int,
    FOREIGN KEY (id_genre) REFERENCES genre(id_genre) ON DELETE CASCADE,
    FOREIGN KEY (id_salon) REFERENCES salon(id_salon) ON DELETE CASCADE
);

create table collection_details(
	id_collection int not null,
    id_rec int not null
);

Alter table collection_details ADD FOREIGN KEY (id_collection) REFERENCES collections(id_collection) ON DELETE CASCADE;
Alter table collection_details ADD FOREIGN KEY (id_rec) REFERENCES records(id_rec) ON DELETE CASCADE;

create table collections_in_salon(
	price int not null,
    sale int,
    count_ int,
    id_collection int not null,
    id_salon int,
    FOREIGN KEY (id_collection) REFERENCES collections(id_collection) ON DELETE CASCADE,
    FOREIGN KEY (id_salon) REFERENCES salon(id_salon) ON DELETE CASCADE
);

create table records_in_salon(
	price int not null,
    sale int,
    count_ int,
    id_rec int not null,
    id_salon int,
    FOREIGN KEY (id_rec) REFERENCES records(id_rec) ON DELETE CASCADE,
    FOREIGN KEY (id_salon) REFERENCES salon(id_salon) ON DELETE CASCADE
);

create table administrator(
	id_admin int not null PRIMARY KEY AUTO_INCREMENT,
    id_user int not null, 
    FOREIGN KEY (id_user) REFERENCES users(id_user) ON DELETE CASCADE
);

create table seller(
	id_seller int not null PRIMARY KEY AUTO_INCREMENT,
    id_user int not null, 
    id_salon int not null,
    FOREIGN KEY (id_user) REFERENCES users(id_user) ON DELETE CASCADE,
    FOREIGN KEY (id_salon) REFERENCES salon(id_salon) ON DELETE CASCADE
);


