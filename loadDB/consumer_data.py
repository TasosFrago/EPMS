from typing import Optional
from random import randrange, random

from pydantic import BaseModel
from faker import Faker


cell_codes = ["685185", "685500", "685505", "685550", "690000", "690002", "690003", "690069", "690345", "694", "695000", "69601", "698", "697", "698", "69601"]
landline_codes = ["21", "231", "241", "251", "261", "271", "281", "291", "222", "233", "234", "235", "237", "238", "239", "242", "243", "244", "246", "247", "248", "249", "252", "253", "254", "255", "256", "257", "259", "262", "263", "264", "265", "266", "267", "268", "269", "272", "273", "274", "275", "276", "277", "279", "282", "283", "284", "285", "286", "289", "222", "232"]


def fake_phone_number(isLandLine=False) -> int:
    fake = Faker("el_GR")
    phone_codes = landline_codes if isLandLine else cell_codes
    phone_c = phone_codes[randrange(0, len(phone_codes))]
    return int(phone_c + fake.msisdn()[(len(phone_c)+3):])

class Consumer_t(BaseModel):
    first_name: str
    last_name: str
    email: str
    cell: int
    landline: Optional[int]
    credit_info: Optional[int]

def getConsumerData() -> Consumer_t:
    fake = Faker("el_GR")
    profile = fake.simple_profile()

    return Consumer_t(
        first_name=fake.first_name(),
        last_name=fake.last_name(),
        email=str(profile['mail']),
        cell=fake_phone_number(),
        landline=fake_phone_number(True) if (random() < 0.3) else None,
        credit_info=None
    )