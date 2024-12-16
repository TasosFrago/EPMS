CREATE DATABASE lab2425omada1_EPMS;
USE lab2425omada1_EPMS;
ALTER DATABASE lab2425omada1_EPMS CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE CONSUMER ( -- 1
	   user_id int not null AUTO_INCREMENT,
	   first_name varchar(50) not null,
	   last_name varchar(50) not null,
	   email varchar(62) not null UNIQUE,
	   password binary(60) not null,
	   cell varchar(10) not null UNIQUE,
	   landline varchar(10) UNIQUE,
	   credit_info int,

	   primary key (user_id)
);

CREATE TABLE PROVIDER ( -- 2
	   name varchar(50) not null,
	   phone varchar(10) not null UNIQUE,
	   email varchar(50) not null UNIQUE,
	   password binary(60) not null,

	   primary key (name)
);

CREATE TABLE EMPLOYEE ( -- 3
	   badge int not null AUTO_INCREMENT,
	   first_name varchar(50) not null,
	   last_name varchar(50) not null,
	   email varchar(62) not null UNIQUE,
	   password binary(60) not null,
	   phone varchar(10) not null UNIQUE,
	   salary float DEFAULT 830,

	   primary key (badge)
);

CREATE TABLE PLAN ( -- 4
	   plan_id int not null AUTO_INCREMENT,
	   type varchar(40) not null,
	   price float not null,
	   name varchar(50),
	   provider varchar(50) not null,
	   month varchar(30) not null,
	   year int not null,
	   duration int not null,
	   
	   primary key (plan_id),
	   foreign key (provider) references PROVIDER(name)
);

CREATE TABLE METER ( -- 5
	   supply_id int not null AUTO_INCREMENT,
	   plan int,
	   status bool DEFAULT 0,
	   kWh int DEFAULT 0,
	   address varchar(100) not null,
	   rated_power int not null,
	   owner int not null,
	   adgent int,

	   primary key (supply_id),
	   foreign key (owner) references CONSUMER(user_id),
	   foreign key (plan) references PLAN(plan_id),
	   foreign key (adgent) references EMPLOYEE(badge)
);

CREATE TABLE INVOICE ( -- 6
	   invoice_id int not null AUTO_INCREMENT,
	   total float DEFAULT 0,
	   current_cost float not null,
	   receiver int not null,
	   meter int not null,
	   provider varchar(50) not null,
	   plan int not null,

	   primary key (invoice_id),
	   foreign key (receiver) references CONSUMER(user_id),
	   foreign key (provider) references PROVIDER(name),
	   foreign key (meter) references METER(supply_id),
	   foreign key (plan) references PLAN(plan_id)
);

CREATE TABLE CHOOSES ( -- 7
	   user int not null,
	   plan int not null,

	   primary key (user, plan),
	   foreign key (user) references CONSUMER(user_id),
	   foreign key (plan) references PLAN(plan_id)
);

CREATE TABLE PAYS ( -- 8
	   payment_id int not null AUTO_INCREMENT,
	   user int not null,
	   provider varchar(50) not null,
	   supply_id int not null,
	   amount float DEFAULT 0,

	   primary key (payment_id),

	   foreign key (user) references CONSUMER(user_id),
	   foreign key (provider) references PROVIDER(name)
);
