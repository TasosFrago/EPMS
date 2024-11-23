from typing import Optional

from pydantic import BaseModel

class Provider_t(BaseModel):
    name: str
    phone: int
    email: str

providers: list[Provider_t] = [
    Provider_t(
        name = "ELPEDISON",
        phone = 2181818128,
        email = "customercare@elpedison.gr"
    ),
    Provider_t(
        name = "EUNICE",
        phone = 2103242020,
        email = "info@eunice-group.com"
    ),
    Provider_t(
        name = "NRG",
        phone = 2188818101,
        email = "cs@nrg.gr"
    ),
    Provider_t(
        name = "OTE ESTATE",
        phone = 2106372700,
        email = "info@ote-estate.gr"
    ),
    Provider_t(
        name = "PROTERGIA",
        phone = 2103448500,
        email = "info@prot-energy.gr"
    ),
    Provider_t(
        name = "SOLAR ENERGY",
        phone = 2385044506,
        email = "supply@soumpasis.gr"
    ),
    Provider_t(
        name = "VOLTERRA",
        phone = 2130883000,
        email = "customercare@volterra.gr"
    ),
    Provider_t(
        name = "VOLTON",
        phone = 2163001000,
        email = "info@volton.gr"
    ),
    Provider_t(
        name = "ΔΕΗ",
        phone = 8009001000,
        email = "info@dei.com.gr"
    ),
    Provider_t(
        name = "ΕΛΙΝΟΙΛ",
        phone = 2106241500,
        email = "info@elin.gr"
    ),
    Provider_t(
        name = "ΖΕΝΙΘ",
        phone = 2311223099,
        email = "info@zenith.gr"
    ),
    Provider_t(
        name = "ΗΡΩΝ",
        phone = 2130075499,
        email = "customercare@heron.gr"
    ),
    Provider_t(
        name = "ΦΥΣΙΚΟ ΑΕΡΙΟ ΕΛΛΗΝΙΚΗ ΕΤΑΙΡΙΑ ΕΝΕΡΓΕΙΑΣ",
        phone = 2177181133,
        email = "customerservice@fysikoaerioellados.gr"
    )
]