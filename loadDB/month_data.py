from typing import Optional

from pydantic import BaseModel

class Month_t(BaseModel):
    name: str
    year: int
    
months: list[Month_t] = [
    Month_t(
        name = "January",
        year = 2024
    ),
    Month_t(
        name = "February",
        year = 2024
    ),
    Month_t(
        name = "March",
        year = 2024
    ),
    Month_t(
        name = "April",
        year = 2024
    ),
    Month_t(
        name = "May",
        year = 2024
    ),
    Month_t(
        name = "June",
        year = 2024
    ),
    Month_t(
        name = "July",
        year = 2024
    ),
    Month_t(
        name = "August",
        year = 2024
    ),
    Month_t(
        name = "September",
        year = 2024
    ),
    Month_t(
        name = "October",
        year = 2024
    ),
    Month_t(
        name = "November",
        year = 2024
    ),
    Month_t(
        name = "December",
        year = 2024
    ),
    Month_t(
        name = "January",
        year = 2025
    ),
    Month_t(
        name = "February",
        year = 2025
    ),
    Month_t(
        name = "March",
        year = 2025
    ),
    Month_t(
        name = "April",
        year = 2025
    ),
    Month_t(
        name = "May",
        year = 2025
    ),
    Month_t(
        name = "June",
        year = 2025
    ),
    Month_t(
        name = "July",
        year = 2025
    ),
    Month_t(
        name = "August",
        year = 2025
    ),
    Month_t(
        name = "September",
        year = 2025
    ),
    Month_t(
        name = "October",
        year = 2025
    ),
    Month_t(
        name = "November",
        year = 2025
    ),
    Month_t(
        name = "December",
        year = 2025
    ),
]