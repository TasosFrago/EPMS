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

CREATE TABLE DEPARTMENT ( -- 3
	region varchar(30) not null,
	phone varchar(10) not null UNIQUE,

	primary key (region)
);

CREATE TABLE EMPLOYEE ( -- 4
	badge int not null AUTO_INCREMENT,
	first_name varchar(50) not null,
	last_name varchar(50) not null,
	email varchar(62) not null UNIQUE,
	password binary(60) not null,
	phone varchar(10) not null UNIQUE,
	salary float DEFAULT 830,
	department varchar(30) not null,

	primary key (badge),

	foreign key (department) references DEPARTMENT(region)
);

CREATE TABLE PLAN ( -- 5
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

CREATE TABLE METER ( -- 6
	supply_id int not null AUTO_INCREMENT,
	plan int,
	status bool DEFAULT 0,
	kWh int DEFAULT 0,
	address varchar(100) not null,
	rated_power int not null,
	owner int not null,
	department varchar(30) not null,
	agent int,

	primary key (supply_id),

	foreign key (plan) references PLAN(plan_id),
	foreign key (owner) references CONSUMER(user_id),
	foreign key (department) references DEPARTMENT(region),
	foreign key (agent) references EMPLOYEE(badge)
);

CREATE TABLE INVOICE ( -- 7
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

-- VIEW for calculating unpaid invoices
CREATE VIEW INVOICE_PAYMENT_STATUS AS
SELECT i.invoice_id,
	   i.receiver as consumer,
       i.total AS invoice_total,
       IFNULL(SUM(p.amount), 0) AS total_paid,
       (IFNULL(SUM(p.amount), 0) >= i.total) AS is_paid
FROM INVOICE i
LEFT JOIN PAYS p
ON i.receiver = p.user AND i.provider = p.provider AND i.meter = p.supply_id
GROUP BY i.invoice_id, i.total
ORDER BY i.invoice_id;

-- VIEW for calculating average meter per employee for each region
CREATE VIEW AVERAGE_METERS_PER_EMPLOYEE AS
SELECT d.region AS department, 
	   COUNT(m.supply_id) / (COUNT(DISTINCT e.badge))
	   AS average_meters_per_employee
FROM DEPARTMENT d
LEFT OUTER JOIN EMPLOYEE e ON d.region = e.department
LEFT OUTER JOIN METER m ON d.region = m.department
GROUP BY d.region;

-- Function to get meter's owner's phone number
/*

\d $$

CREATE FUNCTION get_phone
	(meter INT)
RETURNS VARCHAR(10)
BEGIN
	DECLARE phone VARCHAR(10);
	SET phone = (SELECT cell
				 FROM CONSUMER, METER
				 WHERE user_id = owner AND supply_id = meter);
	RETURN phone;
END$$

\d ;

*/

--Function to find invoice date (DEPLOY AT YOUR OWN RISK)
/*

\d $$

CREATE FUNCTION get_invoice_date
	(invoiceID INT,
	 invoice_plan INT,
	 invoice_meter INT)
RETURNS DATE
BEGIN
	DECLARE month_year DATE;
	DECLARE month_year_string VARCHAR(30);
	DECLARE month_offset INT;

	SET month_offset = (SELECT COUNT(invoice_id)
						FROM INVOICE
						WHERE meter = invoice_meter
							  AND plan = invoice_plan
							  AND invoice_id < invoiceID);

	SET month_year_string = (SELECT CONCAT("1-", month, "-", year)
							 FROM PLAN
							 WHERE plan_id = invoice_plan);

	SET month_year = STR_TO_DATE(month_year_string, "%d-%M-%Y");

	SET month_year = (DATE_ADD(month_year, INTERVAL month_offset MONTH));

	RETURN month_year;
END;

\d ;

*/

-- Procedure to cut power
/*

\d $$

CREATE PROCEDURE cut_power
	(IN meter_id INT)
BEGIN
	DECLARE money_owed FLOAT;

	SET money_owed = ((SELECT SUM(current_cost)
					   FROM INVOICE
					   WHERE meter = meter_id) -
					  (SELECT SUM(amount)
					   FROM PAYS
					   WHERE supply_id = meter_id));

	IF money_owed > 1500
	THEN
		UPDATE METER
		SET status = 0
		WHERE supply_id = meter_id;
	END IF;
END;

\d ;

*/