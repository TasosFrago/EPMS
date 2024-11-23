CREATE DATABASE lab2425omada1_EPMS;
USE lab2425omada1_EPMS;
ALTER DATABASE lab2425omada1_EPMS CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE CONSUMER (
	   user_id int not null AUTO_INCREMENT,
	   first_name varchar(50) not null,
	   last_name varchar(50) not null,
	   email varchar(62) not null UNIQUE,
	   cell varchar(10) not null UNIQUE,
	   landline varchar(10) UNIQUE,
	   credit_info int,

	   primary key (user_id)
);

CREATE TABLE PROVIDER (
	   name varchar(50) not null,
	   phone varchar(10) not null UNIQUE,
	   email varchar(50) not null UNIQUE,

	   primary key (name)
);

CREATE TABLE METER (
	   supply_id int not null AUTO_INCREMENT,
	   status bool not null,
	   kWh int DEFAULT 0,
	   address varchar(100) not null,
	   rated_power int not null,
	   owner int not null,

	   primary key (supply_id),
	   foreign key (owner) references CONSUMER(user_id)
);

CREATE TABLE PLAN (
	   plan_id int not null AUTO_INCREMENT,
	   type varchar(40) not null,
	   price float not null,
	   name varchar(50),
	   provider varchar(50) not null,
	   
	   primary key (plan_id),
	   foreign key (provider) references PROVIDER(name)
);

CREATE TABLE MONTH (
	   name varchar(30) not null,
	   year int not null,

	   primary key (name, year)
);

/*
 *	id int unsigned: bigger number
 *	auto_increament: use LAST_INSERT_ID()
 *	ALTER TABLE tbl AUTO_INCREMENT = num; to start from bigger number
 */
CREATE TABLE INVOICE (
	   invoice_id int not null AUTO_INCREMENT,
	   total int DEFAULT 0,
	   current_cost int not null,
	   receiver int not null,
	   meter int not null,
	   provider varchar(50) not null,
	   month varchar(30) not null,
	   year int not null,

	   primary key (invoice_id),
	   foreign key (receiver) references CONSUMER(user_id),
	   foreign key (provider) references PROVIDER(name),
	   foreign key (meter) references METER(supply_id),
	   foreign key (month, year) references MONTH(name, year)
);

CREATE TABLE AVAILABILITY(
	   year int not null,
	   month varchar(30) not null,
	   plan int not null,

	   primary key (year, month, plan),
	   foreign key (month, year) references MONTH(name, year),
	   foreign key (plan) references PLAN(plan_id)
);

CREATE TABLE CHOOSES (
	   user int not null,
	   plan int not null,
	   supply_id int not null,

	   primary key (user, plan, supply_id),
	   foreign key (user) references CONSUMER(user_id),
	   foreign key (plan) references PLAN(plan_id),
	   foreign key (supply_id) references METER(supply_id)
);

CREATE TABLE PAYS (
	   user int not null,
	   provider varchar(50) not null,
	   supply_id int not null,
	   amount int DEFAULT 0,

	   primary key (user, provider),

	   foreign key (user) references CONSUMER(user_id),
	   foreign key (provider) references PROVIDER(name)
);
