import random

from pydantic import BaseModel
from faker import Faker


def gen_kWh():
    """Generate believable kWh spend per month"""
    highValProb = 0.05 # 5% chance for high spender

    if(random.random() < highValProb):
        return random.randint(5000, 10000)
    else:
        return int(random.triangular(500, 2000, 1000))

class Meter_t(BaseModel):
    status: bool
    kWh: int
    address: str
    rated_power: int
    owner: int

def getMeterData(owner_id: int) -> Meter_t:
    fake = Faker("el_GR")

    return Meter_t(
        status = True, # Assume that account is established so always TRUE
        kWh = 0, # New connection so kWh is reset to 0
        address = fake.line_address(),
        rated_power = 500,
        owner = owner_id
    )