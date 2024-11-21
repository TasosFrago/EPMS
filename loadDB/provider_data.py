from typing import Optional

from pydantic import BaseModel


class Provider_t(BaseModel):
    element: int # attribute: type
    optional_element: Optional[str]


providers: list[Provider_t] = [
    Provider_t(
        element = 2, # example of provider
        optional_element = None # optional attributes need to have value None
    ),
    Provider_t(
        element = 1,
        optional_element = "Test"
    )
]
