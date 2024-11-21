CREATE TABLE CONSUMER (
	   user_id int not null,
	   first_name varchar(50) not null,
	   last_name varchar(50) not null,
	   email varchar(62) not null,
	   cell int not null,
	   landline int,
	   credit_info int,

	   primary key(user_id)
);

CREATE TABLE PROVIDER (
	   name varchar(40) not null,
	   phone int not null,
	   email varchar(50) not null,

	   primary key(name)
);

CREATE TABLE METER (
	   supply_id int not null,
	   status bool not null,
	   kWh int not null,
	   address varchar(100) not null,
	   rated_power int not null,
	   owner int not null,

	   primary key(supply_id)
);

CREATE TABLE PLAN (
	   plan_id int not null,
	   type varchar(40) not null,
	   price int not null,
	   name varchar(50),
	   provider varchar(40) not null,
	   
	   primary key(plan_id)
);

CREATE TABLE MONTH (
	   name varchar(30) not null,
	   year int not null,

	   primary key(name),
	   primary key(year)
);

/*
 *	id int unsigned: bigger number
 *	auto_increament: use LAST_INSERT_ID()
 *	ALTER TABLE tbl AUTO_INCREMENT = num; to start from bigger number
 */
CREATE TABLE INVOICE (
	   invoice_id int not null,
	   total int not null DEFAULT 0,
	   current_cost int not null,
	   receiver int,
	   associated_with int not null,
	   provider varchar(40) not null,
	   month varchar(30) not null,
	   year int not null,

	   primary key(invoice_id)
);

CREATE TABLE AVAILABILITY(
	   year int not null,
	   month varchar(30) not null,
	   plan int not null,

	   primary key(year),  -- TODO 
	   primary key(month), -- TODO 
	   primary key (plan)  -- TODO 
);

CREATE TABLE CHOOSES (
	   user int not null,
	   plan int not null,
	   supply_id int not null,

	   primary key(user),
	   primary key(plan),
	   primary key(supply_id)
);

CREATE TABLE PAYS (
	   user int not null,
	   provider varchar(40) not null,
	   supply_id int not null,
	   amount int not null DEFAULT 0,

	   primary key(user),
	   primary key(provider)
);
