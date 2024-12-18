import random
from typing import Optional

from pydantic import BaseModel
from faker import Faker


def gen_meterNum() -> int:
    """Simulate the probality of owning 1 or more properties(i.e have more meters)"""
    P_more_meters = random.uniform(0, 1) <= 0.3 # 70% chache of owning one

    if P_more_meters:
        return random.randint(2, 3) # Simulate having 2 to 3 meters
    else:
        return 1

def gen_kWh() -> int:
    """Generate believable kWh spend per month"""
    highValProb = 0.05 # 5% chance for high spender
    if(random.random() < highValProb):
        return random.randint(800, 1200)
    else:
        return int(random.triangular(200, 800, 400))

class Meter_t(BaseModel):
    plan: Optional[int]
    status: Optional[int]
    kWh: Optional[int]
    address: str
    rated_power: int
    owner: int
    agent: Optional[int]

def getMeterData(owner_id: int) -> Meter_t:
    fake = Faker("el_GR")

    return Meter_t(
        plan = None,
        status = None,
        kWh = None,
        address = fake.line_address(),
        rated_power = 8,
        owner = owner_id,
        agent = None
    )
