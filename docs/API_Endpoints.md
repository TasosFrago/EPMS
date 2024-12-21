## AUTHENTICATION

To be expanded...

## CONSUMER

⚠ ⚠ ⚠ consumer/user_id reserved for dashboard ⚠ ⚠ ⚠

▶ Account details (GET) /consumer/user_id/account
    · user_id ──> first_name, last_name, email, cell, landline

▶ New account (POST) /consumer
    · first_name, last_name, email, password, cell, landline

▶ Update account (PATCH) /consumer/user_id/account
    · user_id <── first_name, last_name, email, passsword, cell, landline

▶ Delete account (DELETE) /consumer/user_id/account
    · user_id
    · password

▶ Meter list (GET) /consumer/user_id/meters
    · user_id ──> supply_id, status, address

▶ Meter information (GET) /consumer/user_id/meter/supply_id
    · supply_id ──> plan, status, kWh, address, rated_power, department

▶ New meter (POST) /consumer/user_id/meters
    · address, rated_power, department
    · owner <── user_id

▶ Delete meter (DELETE) /consumer/user_id/meters/supply_id
    · supply_id
    · password

▶ Available plans (GET) /consumer/user_id/meters/supply_id/plans
    · supply_id, plan, month ──> plan_id

▶ Choose plan (PATCH) /consumer/user_id/meters/supply_id
    · supply_id <── plan_id

▶ Invoice list (GET) /consumer/user_id/invoices
    · user_id ──> invoice_id, current_cost, provider, plan(month, year)
    · user_id, supply_id ──> invoice_id, current_cost, provider, plan(month, year)

▶ Invoice details (GET) /consumer/user_id/invoices/invoice_id
    · invoice_id ──> current_cost, total, meter, provider, plan(month, year)

▶ Make new payment (POST) /consumer/user_id/pay
    · user_id, provider_name, supply_id, ammount, credit_info
    · user_id, invoice_id, credit_info ──> provider_name, supply_id, ammount

▶ Pay specific invoice (POST) /consumer/user_id/invoices/invoice_id/pay
    · user_id, provider_name, supply_id, ammount, credit_info

▶ Payment history (GET) /consumer/user_id/payments
    · user_id, (sorting) ──> payment_id, supply_id, provider, ammount
    · user_id, supply_id, (sorting) ──> payment_id, provider, ammount
    · user_id, provider, (sorting) ──> payment_id, supply_id, ammount
    · user_id, provider, supply_id, (sorting) ──> payment_id, ammount

## PROVIDER

⚠ ⚠ ⚠ /provider/name reserved for dashboard ⚠ ⚠ ⚠

▶ Account details (GET) /provider/name/account
    · name ──> phone, email

▶ New account (POST) /provider
    · name, phone, email, password

▶ Update account (PATCH) /provider/name/account
    · name <── phone, email, password

▶ Delete account (DELETE) /provider/name/account
    · name
    · password

▶ Meter list (GET) /provider/name/meters
    · name ──> supply_id

▶ Meter information (GET) /provider/name/meters/supply_id
    · supply_id ──> department, address, plan, rated_power, status, kWh, owner

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