import random
from typing import Optional
    
from consumer_data import fake_phone_number
from department_data import departments
from password_data import get_password

from pydantic import BaseModel
from faker import Faker


class Employee_t(BaseModel):
    first_name: str
    last_name: str
    email: str
    password: str
    phone: str
    salary: Optional[float]
    department: str
    
def getEmployeeData() -> Employee_t:
    fake = Faker()
    profile = fake.simple_profile()
    full_name = str(profile["name"]).split()

    first_name = full_name[0]
    last_name = " ".join(full_name[1:])
    
    regions = [dept.region for dept in departments]
    rand_dept = random.choice(regions)

    return Employee_t(
        first_name = first_name,
        last_name = last_name,
        email = str(profile["mail"]),
        password = get_password(str(profile["mail"])),
        phone = fake_phone_number(),
        salary = None,
        department = rand_dept
    )
