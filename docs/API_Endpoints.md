## AUTHENTICATION

To be expanded...

## CONSUMER

⚠ ⚠ ⚠ consumer/user_id reserved for dashboard ⚠ ⚠ ⚠

### Account details 

**(GET)** `/consumer/{user_id}/`
    · user_id ──> first_name, last_name, email, cell, landline
	```json
	{
		"first_name": "example",
		"last_name": "example",
		"email": "example@domain.com",
		"cell": "2544325235",
		"landline": "",
	}
	```

### Create new account 

**(POST)** `/auth/signup/consumer`
    · first_name, last_name, email, password, cell, landline

	```json
	// POST data
	{
		"first_name": "example",
		"last_name": "example",
		"email": "example@domain.com",
		"cell": "2544325235",
		"password": "password123",
		"landline": "",
	}
	```

### Update account 

**(PATCH)** `/consumer/{user_id}/`
    · user_id <── first_name, last_name, email, passsword, cell, landline
	- If any of the data is not given it is not updated

	```json
	// PATCH data
	{
		"first_name": "example",
		"last_name": "example",
		"email": "example@domain.com",
		"cell": "2544325235",
		"password": "password123",
		"landline": "",
	}
	```

### Delete account 

**(DELETE)** `/consumer/{user_id}/`
    · user_id
    · password

### Meter list 

**(GET)** `/consumer/{user_id}/meters`
    · user_id ──> supply_id, status, address
	
	```json
	// GET data
	[
		{
			"supply_id": 121,
			"status": 1,
			"address": "Kavala",
		},
		...
	]
	```

### Meter information 

**(GET)** `/consumer/{user_id}/meter/{supply_id}`
	- supply_id ──> plan, status, kWh, address, rated_power, department
	
	```json
	// GET data
	{
		"plan": 121,
		"status": 1,
		"kWh": 834,
		"address": "Kavala",
		"rated_power": 133,
		"department": "Kavala",
	}
	```

### New meter 

**(POST)** `/consumer/{user_id}/meter`
	- address, rated_power, department
    - owner <── user_id

### Delete meter 

**(DELETE)** `/consumer/{user_id}/meters/{supply_id}`
    - supply_id
    - password

### Available plans 

**(GET)** `/consumer/{user_id}/meters/{supply_id}/plans`
    - supply_id, plan, month ──> plan_id

	```json
	// GET data
	[
		{
			"type": "GREEN",
			"price": 0.23,
			"name": "example",
			"provider": "DEI",
			"month": "JUNE",
			"year": 2024,
			"duration": 1,
		},
		...
	]
	```

### Choose plan 

**(PATCH)** `/consumer/{user_id}/meters/{supply_id}`
    - supply_id <── plan_id
	
	```json
	// PATCH data
	{
		"plan_id": 23,
	}
	```

### Invoice list 

**(GET)** `/consumer/{user_id}/invoices`
    - user_id ──> invoice_id, current_cost, provider, plan(month, year)
    - user_id, supply_id ──> invoice_id, current_cost, provider, plan(month, year)
	
	```json
	// GET data
	[
		{
			"invoice_id": 324,
			"current_cost": 324,
			"provider": "DEI",
			"supply_id": 134,
			"plan_name": "example",
			"month": "JUNE",
			"year": 2025,
		},
		...
	]
	```

### Invoice details

**(GET)** `/consumer/{user_id}/invoices/{invoice_id}`
    - invoice_id ──> current_cost, total, meter, provider, plan(month, year)
	
	```json
	// GET data
	{
		"current_cost": 134.2,
		"total": 1324.2,
		"meter": 314,
		"provider": "DEI",
		"plan_name": "example",
		"plan_type": "GREEN",
		"month": "JUNE",
		"year": 2024,
	}
	```

### Make new payment 

**(POST)** `/consumer/{user_id}/meters/{supply_id}/pay`
    - user_id, provider_name, supply_id, ammount, credit_info
    - user_id, invoice_id, credit_info ──> provider_name, supply_id, ammount
	
	```json
	// POST data
	{
		"provider": "DEI",
		"amount": 234.0,
		"credit_info": 234, // Might not implement
	}
	```

**DON'T NEED THIS, WILL IMPLEMENT IT ON THE CLIENT**
### Pay specific invoice 

**(POST)** `/consumer/{user_id}/invoices/{invoice_id}/pay`
    - user_id, provider_name, supply_id, ammount, credit_info

### Payment history 

**(GET)** `/consumer/{user_id}/payments`
    - user_id, (sorting) ──> payment_id, supply_id, provider, ammount
    - user_id, supply_id, (sorting) ──> payment_id, provider, ammount
    - user_id, provider, (sorting) ──> payment_id, supply_id, ammount
    - user_id, provider, supply_id, (sorting) ──> payment_id, ammount
	
	```json
	// GET data
	[
		{
			"provider": "DEI",
			"supply_id": 234,
			"amount": 234.2,
		},
		...
	]
	```

## PROVIDER

⚠ ⚠ ⚠ /provider/name reserved for dashboard ⚠ ⚠ ⚠

### Account details 

**(GET)** `/provider/{name}/`
    - name ──> phone, email
	
	```json
	// GET data
	{
		"name": "DEI",
		"phone": "2235245324",
		"email": "dei@domain.com",
	}
	```

### New account 

**(POST)** `auth/signup/provider`
    - name, phone, email, password
	
	```json
	// POST data
	{
		"name": "DEI",
		"email": "example@domain.com",
		"phone": "4234625362",
		"passowrd": "password123",
	}
	```

### Update account 

**(PATCH)** `/provider/{name}/`
    - name <── phone, email, password
	
	```json
	// PATCH data
	{
		"name": "DEI",
		"email": "example@domain.com",
		"phone": "46236325",
		"password": "password134",
	}
	```

### Delete account 

**(DELETE)** `/provider/{name}`
    - name
    - password

### Meter list 

**(GET)** `/provider/{name}/meters`
    - name ──> supply_id
	
	```json
	// GET data
	[
		{
			"supply_id": 34,
		},
		...
	]
	```

### Meter information 

**(GET)** `/provider/{name}/meters/{supply_id}`
    - supply_id ──> department, address, plan, rated_power, status, kWh, owner

▶ Contact meter owner (GET) /provider/name/meters/supply_id/owner
    · user_id ──> first_name, last_name, email, cell, landline

▶ Meter invoices (GET) /provider/name/meters/supply_id/invoices
    · supply_id ──> invoice_id, current_cost, total, plan(month, year)

▶ Departments list (GET) /provider/name/departments
    · region, phone

▶ Available plans (GET) /provider/name/meters/supply_id/plans
    · supply_id, plan, month ──> plan_id

▶ Choose plan (PATCH) /provider/name/meters/supply_id
    · supply_id <── plan_id

▶ Plans list (GET) /provider/name/plans
    · name ──> type, plan_name, price, month, year, duration
    · name, type ──> plan_name, price, month, year, duration
    · name, month ──> type, plan_name, price, year, duration

▶ Plan information (GET) /provider/name/plans/plan_id
    · plan_id ──> type, price, plan_name, month, year, duration, count(supply_id)

▶ Create new plan (POST) /provider/name/plans/new
    · provider, type, plan_name, price, month, year, duration

▶ Delete plan (DELETE) /provider/name/plans/plan_id
    · plan_id
    · password

▶ List of invoices (GET) /provider/name/invoices
    · name ──> invoice_id, current_cost, total, plan(month, year)

▶ Invoice details (GET) /provider/name/invoices/invoice_id
    · invoice_id ──> receiver, meter, plan(month, year), current_cost, total

▶ Issue invoice (POST) /provider/name/meters/supply_id/invoices/issue (if kWh > 0)
    · supply_id, receiver, provider, plan(month, year), current_cost, total

▶ List of payments (GET) /provider/name/payments
    · name ──> payment_id, supply_id, user, ammount
    · name, supply_id ──> payment_id, user, ammount
    · name, user_id ──> payment_id, supply_id, ammount

## DEPARTMENT

⚠ ⚠ ⚠ /department/region reserved for dashboard ⚠ ⚠ ⚠

No login needed for departments. Assume they use their own database.

▶ Update phone number (PATCH) /department/region/phone
    · region <── phone

▶ New employee (POST) /department/region/employees/new
    · first_name, last_name, email, password, phone, department

▶ Employee list (GET) /department/region/employees
    · region ──> badge, first_name, last_name

▶ Employee details (GET) /department/region/employees/badge
    · badge ──> first_name, last_name, email, phone, salary

▶ Update employee salary (PATCH) /department/region/employees/badge/salary
    · badge <── salary

▶ Delete employee (DELETE) /department/region/employees/badge
    · badge

▶ Meter list (GET) /department/region/meters
    · region ──> supply_id, agent, status, address
    · region, status ──> supply_id, agent, address
    · region, agent ──> supply_id, status, address
    · region, agent, status ──> supply_id, address

▶ Meter information (GET) /department/region/meters/supply_id
    · supply_id ──> agent, address, status, kWh, rated_power, owner

▶ Contact meter owner (GET) /department/region/meters/supply_id/owner
    · user_id ──> first_name, last_name, email, cell, landline

▶ Provider list (GET) /department/region/providers
    · provider_name

▶ Contact provider (GET) /department/region/providers/name
    · provider_name ──> phone, email

## EMPLOYEE

⚠ ⚠ ⚠ employee/badge reserved for dashboard ⚠ ⚠ ⚠

▶ Account details (GET) /employee/badge/account
    · badge ──> first_name, last_name, email, phone, salary, department

▶ Update account (PATCH) /employee/badge/account
    · badge <── first_name, last_name, email, password, phone

▶ Contact office (GET) /employee/badge/office
    · badge ──> department_phone

▶ List of assigned meters (GET) /employee/badge/meters
    · badge ──> supply_id, address, status
    · badge, status ──> supply_id, address

▶ Meter information (GET) /employee/badge/meters/supply_id
    · supply_id ──> address, status, rated_power, owner

▶ Contact meter owner (GET) /employee/badge/meters/supply_id/owner
    · user_id ──> first_name, last_name, email, cell, landline

▶ Insert meter reading (PATCH) /employee/badge/meters/supply_id
    · supply_id <── kWh

▶ Change meter status (PATCH) /emplyee/badge/meters/supply_id
    · supply_id <── status
