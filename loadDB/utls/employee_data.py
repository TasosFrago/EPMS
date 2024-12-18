from typing import Optional
    
from utls.consumer_data import fake_phone_number
from utls.password_data import get_password

from pydantic import BaseModel
from faker import Faker


class Employee_t(BaseModel):
    first_name: str
    last_name: str
    email: str
    password: str
    phone: str
    salary: Optional[float]

    
def getEmployeeData() -> Employee_t:
    fake = Faker()
    profile = fake.simple_profile()
    full_name = str(profile["name"]).split()

    first_name = full_name[0]
    last_name = " ".join(full_name[1:])

    return Employee_t(
        first_name = first_name,
        last_name = last_name,
        email = str(profile["mail"]),
        password = get_password(str(profile["mail"])),
        phone = fake_phone_number(),
        salary = None
    )