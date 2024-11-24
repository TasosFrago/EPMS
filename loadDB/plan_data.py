from pydantic import BaseModel, field_validator

class Plan_t(BaseModel):
    name: str
    type: str       # YELLOW, BLUE, GREEN
    price: float
    provider: str
    month: int      # 1-24
    duration: int
    
    @field_validator("type")
    @classmethod
    def validate_type(cls, v: str):
        valid_types = [
            "GREEN",
            "YELLOW",
            "BLUE"
        ] 
        if v not in valid_types:
            raise ValueError(f"Invalid type: {v}. Must be YELLOW, BLUE or GREEN.")
        return v
    
    @field_validator("provider")
    @classmethod
    def validate_provider(cls, v: str):
        valid_providers = [
            "ELPEDISON",
            "EUNICE",
            "NRG",
            "OTE ESTATE",
            "PROTERGIA",
            "SOLAR ENERGY",
            "VOLTERRA",
            "VOLTON",
            "ΔΕΗ",
            "ΕΛΙΝΟΙΛ",
            "ΖΕΝΙΘ",
            "ΗΡΩΝ",
            "ΦΥΣΙΚΟ ΑΕΡΙΟ ΕΛΛΗΝΙΚΗ ΕΤΑΙΡΙΑ ΕΝΕΡΓΕΙΑΣ"
        ]
        if v not in valid_providers:
            raise ValueError(f"Invalid provider: {v}. Must be one of {valid_providers}.")
        return v
    
    @field_validator("price")
    @classmethod
    def validate_price(cls, v: int):
        if not (0.05 < v < 0.5):
            raise ValueError(f"Invalid price {v}. Must be between 0.05 and 0.5.")
        return v

plans: list[Plan_t] = [
    Plan_t(
        name = "ELPEDISON SMART",
        type = "YELLOW",
        price = 0.13508,
        provider = "ELPEDISON",
        month = 1,
        duration = 1
    ),
    Plan_t(
        name = "Bright",
        type = "BLUE",
        price = 0.169,
        provider = "ELPEDISON",
        month = 1,
        duration = 6
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.17058,
        provider = "ELPEDISON",
        month = 1,
        duration = 1
    ),
    Plan_t(
        name = "We Home",
        type = "YELLOW",
        price = 0.13788,
        provider = "EUNICE",
        month = 1,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.16606,
        provider = "EUNICE",
        month = 1,
        duration = 1
    ),
    Plan_t(
        name = "prime4U",
        type = "YELLOW",
        price = 0.15811,
        provider = "NRG",
        month = 1,
        duration = 1
    ),
    Plan_t(
        name = "nrg fixed 4U",
        type = "BLUE",
        price = 0.151,
        provider = "NRG",
        month = 1,
        duration = 3
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.1412,
        provider = "NRG",
        month = 1,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.198,
        provider = "OTE ESTATE",
        month = 1,
        duration = 1
    ),
    Plan_t(
        name = "Value Fair",
        type = "YELLOW",
        price = 0.1328,
        provider = "PROTERGIA",
        month = 1,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.1426,
        provider = "PROTERGIA",
        month = 1,
        duration = 1
    ),
    Plan_t(
        name = "Protergia Οικιακό Value Safe",
        type = "BLUE",
        price = 0.0975,
        provider = "PROTERGIA",
        month = 1,
        duration = 6
    ),
    Plan_t(
        name = "Οικιακό FIXED",
        type = "BLUE",
        price = 0.36,
        provider = "SOLAR ENERGY",
        month = 1,
        duration = 6
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.1688,
        provider = "SOLAR ENERGY",
        month = 1,
        duration = 1
    ),
    Plan_t(
        name = "Οικιακό S1",
        type = "YELLOW",
        price = 0.15327,
        provider = "SOLAR ENERGY",
        month = 1,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.144,
        provider = "VOLTERRA",
        month = 1,
        duration = 1
    ),
    Plan_t(
        name = "Forward για το Σπίτι",
        type = "YELLOW",
        price = 0.15797,
        provider = "VOLTERRA",
        month = 1,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.14425,
        provider = "VOLTON",
        month = 1,
        duration = 1
    ),
    Plan_t(
        name = "VOLTON YELLOW SIMPLE",
        type = "YELLOW",
        price = 0.14364,
        provider = "VOLTON",
        month = 1,
        duration = 1
    ),
    Plan_t(
        name = "VOLTON BLUE FIXED",
        type = "BLUE",
        price = 0.1575,
        provider = "VOLTON",
        month = 1,
        duration = 3
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.13635,
        provider = "ΔΕΗ",
        month = 1,
        duration = 1
    ),
    Plan_t(
        name = "MyHome 4All",
        type = "YELLOW",
        price = 0.12343,
        provider = "ΔΕΗ",
        month = 1,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.14038,
        provider = "ΕΛΙΝΟΙΛ",
        month = 1,
        duration = 1
    ),
    Plan_t(
        name = "Power On! Home",
        type = "YELLOW",
        price = 0.14856,
        provider = "ΕΛΙΝΟΙΛ",
        month = 1,
        duration = 1
    ),
    Plan_t(
        name = "Power On! Blue",
        type = "BLUE",
        price = 0.149,
        provider = "ΕΛΙΝΟΙΛ",
        month = 1,
        duration = 3
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.16947,
        provider = "ΖΕΝΙΘ",
        month = 1,
        duration = 1
    ),
    Plan_t(
        name = "Power Home Choice",
        type = "YELLOW",
        price = 0.1228,
        provider = "ΖΕΝΙΘ",
        month = 1,
        duration = 1
    ),
    Plan_t(
        name = "Power Home Fixed SE",
        type = "BLUE",
        price = 0.16,
        provider = "ΖΕΝΙΘ",
        month = 1,
        duration = 6
    ),
    Plan_t(
        name = "BASIC HOME",
        type = "GREEN",
        price = 0.14052,
        provider = "ΗΡΩΝ",
        month = 1,
        duration = 1
    ),
    Plan_t(
        name = "SIMPLY GENEROUS HOME",
        type = "YELLOW",
        price = 0.15383,
        provider = "ΗΡΩΝ",
        month = 1,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.14265,
        provider = "ΦΥΣΙΚΟ ΑΕΡΙΟ ΕΛΛΗΝΙΚΗ ΕΤΑΙΡΙΑ ΕΝΕΡΓΕΙΑΣ",
        month = 1,
        duration = 1
    ),
    Plan_t(
        name = "MAXI Home",
        type = "YELLOW",
        price = 0.135,
        provider = "ΦΥΣΙΚΟ ΑΕΡΙΟ ΕΛΛΗΝΙΚΗ ΕΤΑΙΡΙΑ ΕΝΕΡΓΕΙΑΣ",
        month = 1,
        duration = 1
    ),
    Plan_t(
        name = "Home Fixed",
        type = "BLUE",
        price = 0.199,
        provider = "ΦΥΣΙΚΟ ΑΕΡΙΟ ΕΛΛΗΝΙΚΗ ΕΤΑΙΡΙΑ ΕΝΕΡΓΕΙΑΣ",
        month = 1,
        duration = 3
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.1465,
        provider = "ELPEDISON",
        month = 2,
        duration = 1
    ),
    Plan_t(
        name = "ELPEDISON SMART",
        type = "YELLOW",
        price = 0.11084,
        provider = "ELPEDISON",
        month = 2,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.1455,
        provider = "EUNICE",
        month = 2,
        duration = 1
    ),
    Plan_t(
        name = "We Home",
        type = "YELLOW",
        price = 0.11614,
        provider = "EUNICE",
        month = 2,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.12453,
        provider = "NRG",
        month = 2,
        duration = 1
    ),
    Plan_t(
        name = "nrg fixed 4U",
        type = "BLUE",
        price = 0.151,
        provider = "NRG",
        month = 2,
        duration = 3
    ),
    Plan_t(
        name = "prime 4U",
        type = "YELLOW",
        price = 0.12617,
        provider = "NRG",
        month = 2,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.1782,
        provider = "OTE ESTATE",
        month = 2,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.14,
        provider = "PROTERGIA",
        month = 2,
        duration = 1
    ),
    Plan_t(
        name = "Value Simple",
        type = "YELLOW",
        price = 0.1298,
        provider = "PROTERGIA",
        month = 2,
        duration = 1
    ),
    Plan_t(
        name = "Protergia Οικιακό - Value Safe",
        type = "BLUE",
        price = 0.0975,
        provider = "PROTERGIA",
        month = 2,
        duration = 6
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.13975,
        provider = "SOLAR ENERGY",
        month = 2,
        duration = 1
    ),
    Plan_t(
        name = "Οικιακό S1",
        type = "YELLOW",
        price = 0.13019,
        provider = "SOLAR ENERGY",
        month = 2,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.135,
        provider = "VOLTERRA",
        month = 2,
        duration = 1
    ),
    Plan_t(
        name = "Forward για το Σπίτι",
        type = "YELLOW",
        price = 0.13486,
        provider = "VOLTERRA",
        month = 2,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.12476,
        provider = "VOLTON",
        month = 2,
        duration = 1
    ),
    Plan_t(
        name = "VOLTON YELLOW SIMPLE",
        type = "YELLOW",
        price = 0.12014,
        provider = "VOLTON",
        month = 2,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.128,
        provider = "ΔΕΗ",
        month = 2,
        duration = 1
    ),
    Plan_t(
        name = "MyHome 4All",
        type = "YELLOW",
        price = 0.1209,
        provider = "ΔΕΗ",
        month = 2,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.11798,
        provider = "ΕΛΙΝΟΙΛ",
        month = 2,
        duration = 1
    ),
    Plan_t(
        name = "Power On! Home",
        type = "YELLOW",
        price = 0.12566,
        provider = "ΕΛΙΝΟΙΛ",
        month = 2,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.14468,
        provider = "ΖΕΝΙΘ",
        month = 2,
        duration = 1
    ),
    Plan_t(
        name = "Power Home Choice",
        type = "YELLOW",
        price = 0.10054,
        provider = "ΖΕΝΙΘ",
        month = 2,
        duration = 1
    ),
    Plan_t(
        name = "BASIC HOME",
        type = "GREEN",
        price = 0.13924,
        provider = "ΗΡΩΝ",
        month = 2,
        duration = 1
    ),
    Plan_t(
        name = "SIMPLY GENEROUS HOME",
        type = "YELLOW",
        price = 0.12937,
        provider = "ΗΡΩΝ",
        month = 2,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.12688,
        provider = "ΦΥΣΙΚΟ ΑΕΡΙΟ ΕΛΛΗΝΙΚΗ ΕΤΑΙΡΙΑ ΕΝΕΡΓΕΙΑΣ",
        month = 2,
        duration = 1
    ),
    Plan_t(
        name = "MAXI Home",
        type = "YELLOW",
        price = 0.1081,
        provider = "ΦΥΣΙΚΟ ΑΕΡΙΟ ΕΛΛΗΝΙΚΗ ΕΤΑΙΡΙΑ ΕΝΕΡΓΕΙΑΣ",
        month = 2,
        duration = 1
    ),
    Plan_t(
        name = "Home Fixed",
        type = "BLUE",
        price = 0.199,
        provider = "ΦΥΣΙΚΟ ΑΕΡΙΟ ΕΛΛΗΝΙΚΗ ΕΤΑΙΡΙΑ ΕΝΕΡΓΕΙΑΣ",
        month = 2,
        duration = 3
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.1076,
        provider = "ELPEDISON",
        month = 3,
        duration = 1
    ),
    Plan_t(
        name = "ELPEDISON SMART",
        type = "YELLOW",
        price = 0.09821,
        provider = "ELPEDISON",
        month = 3,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.1123,
        provider = "EUNICE",
        month = 3,
        duration = 1
    ),
    Plan_t(
        name = "We Home",
        type = "YELLOW",
        price = 0.1093,
        provider = "EUNICE",
        month = 3,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.115,
        provider = "NRG",
        month = 3,
        duration = 1
    ),
    Plan_t(
        name = "nrg fixed 4U",
        type = "BLUE",
        price = 0.151,
        provider = "NRG",
        month = 3,
        duration = 3
    ),
    Plan_t(
        name = "prime 4U",
        type = "YELLOW",
        price = 0.11693,
        provider = "NRG",
        month = 3,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.139,
        provider = "OTE ESTATE",
        month = 3,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.10776,
        provider = "PROTERGIA",
        month = 3,
        duration = 1
    ),
    Plan_t(
        name = "Value Fair",
        type = "YELLOW",
        price = 0.09176,
        provider = "PROTERGIA",
        month = 3,
        duration = 1
    ),
    Plan_t(
        name = "Protergia Οικιακό - Value Safe",
        type = "BLUE",
        price = 0.095,
        provider = "PROTERGIA",
        month = 3,
        duration = 6
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.1027,
        provider = "SOLAR ENERGY",
        month = 3,
        duration = 1
    ),
    Plan_t(
        name = "Οικιακό S1",
        type = "YELLOW",
        price = 0.12295,
        provider = "SOLAR ENERGY",
        month = 3,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.115,
        provider = "VOLTERRA",
        month = 3,
        duration = 1
    ),
    Plan_t(
        name = "Forward για το Σπίτι",
        type = "YELLOW",
        price = 0.12884,
        provider = "VOLTERRA",
        month = 3,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.10306,
        provider = "VOLTON",
        month = 3,
        duration = 1
    ),
    Plan_t(
        name = "VOLTON YELLOW SIMPLE",
        type = "YELLOW",
        price = 0.11317,
        provider = "VOLTON",
        month = 3,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.10827,
        provider = "ΔΕΗ",
        month = 3,
        duration = 1
    ),
    Plan_t(
        name = "MyHome 4All",
        type = "YELLOW",
        price = 0.09523,
        provider = "ΔΕΗ",
        month = 3,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.08182,
        provider = "ΕΛΙΝΟΙΛ",
        month = 3,
        duration = 1
    ),
    Plan_t(
        name = "Power On! Home",
        type = "YELLOW",
        price = 0.1185,
        provider = "ΕΛΙΝΟΙΛ",
        month = 3,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.10467,
        provider = "ΖΕΝΙΘ",
        month = 3,
        duration = 1
    ),
    Plan_t(
        name = "Power Home Choice",
        type = "YELLOW",
        price = 0.09393,
        provider = "ΖΕΝΙΘ",
        month = 3,
        duration = 1
    ),
    Plan_t(
        name = "BASIC HOME",
        type = "GREEN",
        price = 0.10304,
        provider = "ΗΡΩΝ",
        month = 3,
        duration = 1
    ),
    Plan_t(
        name = "SIMPLY GENEROUS HOME",
        type = "YELLOW",
        price = 0.12173,
        provider = "ΗΡΩΝ",
        month = 3,
        duration = 1
    ),
    Plan_t(
        name = "Fix Genius Home 3",
        type = "BLUE",
        price = 0.0992,
        provider = "ΗΡΩΝ",
        month = 3,
        duration = 3
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.11165,
        provider = "ΦΥΣΙΚΟ ΑΕΡΙΟ ΕΛΛΗΝΙΚΗ ΕΤΑΙΡΙΑ ΕΝΕΡΓΕΙΑΣ",
        month = 3,
        duration = 1
    ),
    Plan_t(
        name = "Ρεύμα MAXI Home 10",
        type = "YELLOW",
        price = 0.0913,
        provider = "ΦΥΣΙΚΟ ΑΕΡΙΟ ΕΛΛΗΝΙΚΗ ΕΤΑΙΡΙΑ ΕΝΕΡΓΕΙΑΣ",
        month = 3,
        duration = 1
    ),
    Plan_t(
        name = "Home Fixed",
        type = "BLUE",
        price = 0.199,
        provider = "ΦΥΣΙΚΟ ΑΕΡΙΟ ΕΛΛΗΝΙΚΗ ΕΤΑΙΡΙΑ ΕΝΕΡΓΕΙΑΣ",
        month = 3,
        duration = 3
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.11705,
        provider = "ELPEDISON",
        month = 4,
        duration = 1
    ),
    Plan_t(
        name = "ELPEDISON SMART UP",
        type = "YELLOW",
        price = 0.085,
        provider = "ELPEDISON",
        month = 4,
        duration = 1
    ),
    Plan_t(
        name = "ELPEDISON Bright Up",
        type = "BLUE",
        price = 0.089,
        provider = "ELPEDISON",
        month = 4,
        duration = 6
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.13,
        provider = "EUNICE",
        month = 4,
        duration = 1
    ),
    Plan_t(
        name = "We Home",
        type = "YELLOW",
        price = 0.10102,
        provider = "EUNICE",
        month = 4,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.10784,
        provider = "NRG",
        month = 4,
        duration = 1
    ),
    Plan_t(
        name = "nrg fixed 4U",
        type = "BLUE",
        price = 0.151,
        provider = "NRG",
        month = 4,
        duration = 3
    ),
    Plan_t(
        name = "prime 4U",
        type = "YELLOW",
        price = 0.11489,
        provider = "NRG",
        month = 4,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.148,
        provider = "OTE ESTATE",
        month = 4,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.11624,
        provider = "PROTERGIA",
        month = 4,
        duration = 1
    ),
    Plan_t(
        name = "Value Simple",
        type = "YELLOW",
        price = 0.09965,
        provider = "PROTERGIA",
        month = 4,
        duration = 1
    ),
    Plan_t(
        name = "Protergia Οικιακό - Value Safe",
        type = "BLUE",
        price = 0.095,
        provider = "PROTERGIA",
        month = 4,
        duration = 6
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.10559,
        provider = "SOLAR ENERGY",
        month = 4,
        duration = 1
    ),
    Plan_t(
        name = "Οικιακό S1",
        type = "YELLOW",
        price = 0.11403,
        provider = "SOLAR ENERGY",
        month = 4,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.099,
        provider = "VOLTERRA",
        month = 4,
        duration = 1
    ),
    Plan_t(
        name = "Forward για το Σπίτι",
        type = "YELLOW",
        price = 0.11874,
        provider = "VOLTERRA",
        month = 4,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.1054,
        provider = "VOLTON",
        month = 4,
        duration = 1
    ),
    Plan_t(
        name = "VOLTON YELLOW SIMPLE",
        type = "YELLOW",
        price = 0.10237,
        provider = "VOLTON",
        month = 4,
        duration = 1
    ),
    Plan_t(
        name = "VOLTON BLUE FIXED 6M",
        type = "BLUE",
        price = 0.088,
        provider = "VOLTON",
        month = 4,
        duration = 6
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.10861,
        provider = "ΔΕΗ",
        month = 4,
        duration = 1
    ),
    Plan_t(
        name = "MyHome 4All",
        type = "YELLOW",
        price = 0.088,
        provider = "ΔΕΗ",
        month = 4,
        duration = 1
    ),
    Plan_t(
        name = "My Home Online",
        type = "BLUE",
        price = 0.142,
        provider = "ΔΕΗ",
        month = 4,
        duration = 3
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.0906,
        provider = "ΕΛΙΝΟΙΛ",
        month = 4,
        duration = 1
    ),
    Plan_t(
        name = "Power On! Home",
        type = "YELLOW",
        price = 0.10968,
        provider = "ΕΛΙΝΟΙΛ",
        month = 4,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.11438,
        provider = "ΖΕΝΙΘ",
        month = 4,
        duration = 1
    ),
    Plan_t(
        name = "Power Home Choice",
        type = "YELLOW",
        price = 0.0837,
        provider = "ΖΕΝΙΘ",
        month = 4,
        duration = 1
    ),
    Plan_t(
        name = "BASIC HOME",
        type = "GREEN",
        price = 0.11555,
        provider = "ΗΡΩΝ",
        month = 4,
        duration = 1
    ),
    Plan_t(
        name = "SIMPLY GENEROUS HOME",
        type = "YELLOW",
        price = 0.11236,
        provider = "ΗΡΩΝ",
        month = 4,
        duration = 1
    ),
    Plan_t(
        name = "Fix Genius Home 3",
        type = "BLUE",
        price = 0.0992,
        provider = "ΗΡΩΝ",
        month = 4,
        duration = 3
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.10604,
        provider = "ΦΥΣΙΚΟ ΑΕΡΙΟ ΕΛΛΗΝΙΚΗ ΕΤΑΙΡΙΑ ΕΝΕΡΓΕΙΑΣ",
        month = 4,
        duration = 1
    ),
    Plan_t(
        name = "Ρεύμα MAXI Home 10",
        type = "YELLOW",
        price = 0.08776,
        provider = "ΦΥΣΙΚΟ ΑΕΡΙΟ ΕΛΛΗΝΙΚΗ ΕΤΑΙΡΙΑ ΕΝΕΡΓΕΙΑΣ",
        month = 4,
        duration = 1
    ),
    Plan_t(
        name = "Home Fixed",
        type = "BLUE",
        price = 0.199,
        provider = "ΦΥΣΙΚΟ ΑΕΡΙΟ ΕΛΛΗΝΙΚΗ ΕΤΑΙΡΙΑ ΕΝΕΡΓΕΙΑΣ",
        month = 4,
        duration = 3
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.10568,
        provider = "ELPEDISON",
        month = 5,
        duration = 1
    ),
    Plan_t(
        name = "ELPEDISON SMART UP",
        type = "YELLOW",
        price = 0.085,
        provider = "ELPEDISON",
        month = 5,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.13,
        provider = "EUNICE",
        month = 5,
        duration = 1
    ),
    Plan_t(
        name = "We Home",
        type = "YELLOW",
        price = 0.12451,
        provider = "EUNICE",
        month = 5,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.10961,
        provider = "NRG",
        month = 5,
        duration = 1
    ),
    Plan_t(
        name = "nrg fixed 4U",
        type = "BLUE",
        price = 0.151,
        provider = "NRG",
        month = 5,
        duration = 3
    ),
    Plan_t(
        name = "prime 4U",
        type = "YELLOW",
        price = 0.13751,
        provider = "NRG",
        month = 5,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.16513,
        provider = "OTE ESTATE",
        month = 5,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.10601,
        provider = "PROTERGIA",
        month = 5,
        duration = 1
    ),
    Plan_t(
        name = "Value Fair",
        type = "YELLOW",
        price = 0.09001,
        provider = "PROTERGIA",
        month = 5,
        duration = 1
    ),
    Plan_t(
        name = "Protergia Οικιακό - Value Safe",
        type = "BLUE",
        price = 0.095,
        provider = "PROTERGIA",
        month = 5,
        duration = 6
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.10085,
        provider = "SOLAR ENERGY",
        month = 5,
        duration = 1
    ),
    Plan_t(
        name = "Οικιακό S1",
        type = "YELLOW",
        price = 0.13901,
        provider = "SOLAR ENERGY",
        month = 5,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.099,
        provider = "VOLTERRA",
        month = 5,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.10168,
        provider = "VOLTON",
        month = 5,
        duration = 1
    ),
    Plan_t(
        name = "VOLTON YELLOW SIMPLE",
        type = "YELLOW",
        price = 0.12511,
        provider = "VOLTON",
        month = 5,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.10656,
        provider = "ΔΕΗ",
        month = 5,
        duration = 1
    ),
    Plan_t(
        name = "MyHome 4All",
        type = "YELLOW",
        price = 0.09353,
        provider = "ΔΕΗ",
        month = 5,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.095,
        provider = "ΕΛΙΝΟΙΛ",
        month = 5,
        duration = 1
    ),
    Plan_t(
        name = "Power On! Home",
        type = "YELLOW",
        price = 0.13445,
        provider = "ΕΛΙΝΟΙΛ",
        month = 5,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.10267,
        provider = "ΖΕΝΙΘ",
        month = 5,
        duration = 1
    ),
    Plan_t(
        name = "Power Home Choice",
        type = "YELLOW",
        price = 0.10525,
        provider = "ΖΕΝΙΘ",
        month = 5,
        duration = 1
    ),
    Plan_t(
        name = "BASIC HOME",
        type = "GREEN",
        price = 0.10118,
        provider = "ΗΡΩΝ",
        month = 5,
        duration = 1
    ),
    Plan_t(
        name = "SIMPLY GENEROUS HOME",
        type = "YELLOW",
        price = 0.13879,
        provider = "ΗΡΩΝ",
        month = 5,
        duration = 1
    ),
    Plan_t(
        name = "Fix Genius Home 3",
        type = "BLUE",
        price = 0.0992,
        provider = "ΗΡΩΝ",
        month = 5,
        duration = 3
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.10954,
        provider = "ΦΥΣΙΚΟ ΑΕΡΙΟ ΕΛΛΗΝΙΚΗ ΕΤΑΙΡΙΑ ΕΝΕΡΓΕΙΑΣ",
        month = 5,
        duration = 1
    ),
    Plan_t(
        name = "Ρεύμα MAXI Home 10",
        type = "YELLOW",
        price = 0.1118,
        provider = "ΦΥΣΙΚΟ ΑΕΡΙΟ ΕΛΛΗΝΙΚΗ ΕΤΑΙΡΙΑ ΕΝΕΡΓΕΙΑΣ",
        month = 5,
        duration = 1
    ),
    Plan_t(
        name = "Home Fixed",
        type = "BLUE",
        price = 0.199,
        provider = "ΦΥΣΙΚΟ ΑΕΡΙΟ ΕΛΛΗΝΙΚΗ ΕΤΑΙΡΙΑ ΕΝΕΡΓΕΙΑΣ",
        month = 5,
        duration = 3
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.17984,
        provider = "ELPEDISON",
        month = 6,
        duration = 1
    ),
    Plan_t(
        name = "ELPEDISON Win",
        type = "YELLOW",
        price = 0.1108,
        provider = "ELPEDISON",
        month = 6,
        duration = 1
    ),
    Plan_t(
        name = "ELPEDISON Bright Up",
        type = "BLUE",
        price = 0.089,
        provider = "ELPEDISON",
        month = 6,
        duration = 6
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.1659,
        provider = "EUNICE",
        month = 6,
        duration = 1
    ),
    Plan_t(
        name = "We Home",
        type = "YELLOW",
        price = 0.14446,
        provider = "EUNICE",
        month = 6,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.15952,
        provider = "NRG",
        month = 6,
        duration = 1
    ),
    Plan_t(
        name = "prime 4U",
        type = "YELLOW",
        price = 0.16064,
        provider = "NRG",
        month = 6,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.192,
        provider = "OTE ESTATE",
        month = 6,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.16422,
        provider = "PROTERGIA",
        month = 6,
        duration = 1
    ),
    Plan_t(
        name = "Value Simple",
        type = "YELLOW",
        price = 0.09093,
        provider = "PROTERGIA",
        month = 6,
        duration = 1
    ),
    Plan_t(
        name = "Protergia Οικιακό - Value Safe",
        type = "BLUE",
        price = 0.095,
        provider = "PROTERGIA",
        month = 6,
        duration = 6
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.12591,
        provider = "SOLAR ENERGY",
        month = 6,
        duration = 1
    ),
    Plan_t(
        name = "Οικιακό S1",
        type = "YELLOW",
        price = 0.16021,
        provider = "SOLAR ENERGY",
        month = 6,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.129,
        provider = "VOLTERRA",
        month = 6,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.12916,
        provider = "VOLTON",
        month = 6,
        duration = 1
    ),
    Plan_t(
        name = "VOLTON YELLOW SIMPLE",
        type = "YELLOW",
        price = 0.14837,
        provider = "VOLTON",
        month = 6,
        duration = 1
    ),
    Plan_t(
        name = "VOLTON BLUE FIXED 6M",
        type = "BLUE",
        price = 0.088,
        provider = "VOLTON",
        month = 6,
        duration = 6
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.11898,
        provider = "ΔΕΗ",
        month = 6,
        duration = 1
    ),
    Plan_t(
        name = "MyHome 4All",
        type = "YELLOW",
        price = 0.11151,
        provider = "ΔΕΗ",
        month = 6,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.1485,
        provider = "ΕΛΙΝΟΙΛ",
        month = 6,
        duration = 1
    ),
    Plan_t(
        name = "Power On! Home",
        type = "YELLOW",
        price = 0.15547,
        provider = "ΕΛΙΝΟΙΛ",
        month = 6,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.16927,
        provider = "ΖΕΝΙΘ",
        month = 6,
        duration = 1
    ),
    Plan_t(
        name = "Power Home Go Electric Plus",
        type = "YELLOW",
        price = 0.11227,
        provider = "ΖΕΝΙΘ",
        month = 6,
        duration = 1
    ),
    Plan_t(
        name = "BASIC HOME",
        type = "GREEN",
        price = 0.16333,
        provider = "ΗΡΩΝ",
        month = 6,
        duration = 1
    ),
    Plan_t(
        name = "SIMPLY GENEROUS HOME",
        type = "YELLOW",
        price = 0.16123,
        provider = "ΗΡΩΝ",
        month = 6,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.12854,
        provider = "ΦΥΣΙΚΟ ΑΕΡΙΟ ΕΛΛΗΝΙΚΗ ΕΤΑΙΡΙΑ ΕΝΕΡΓΕΙΑΣ",
        month = 6,
        duration = 1
    ),
    Plan_t(
        name = "Ρεύμα MAXI Home 10",
        type = "YELLOW",
        price = 0.13052,
        provider = "ΦΥΣΙΚΟ ΑΕΡΙΟ ΕΛΛΗΝΙΚΗ ΕΤΑΙΡΙΑ ΕΝΕΡΓΕΙΑΣ",
        month = 6,
        duration = 1
    ),
    Plan_t(
        name = "Home Fixed",
        type = "BLUE",
        price = 0.199,
        provider = "ΦΥΣΙΚΟ ΑΕΡΙΟ ΕΛΛΗΝΙΚΗ ΕΤΑΙΡΙΑ ΕΝΕΡΓΕΙΑΣ",
        month = 6,
        duration = 3
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.18874,
        provider = "ELPEDISON",
        month = 7,
        duration = 1
    ),
    Plan_t(
        name = "ELPEDISON Win More",
        type = "YELLOW",
        price = 0.0895,
        provider = "ELPEDISON",
        month = 7,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.1823,
        provider = "EUNICE",
        month = 7,
        duration = 1
    ),
    Plan_t(
        name = "We Home",
        type = "YELLOW",
        price = 0.1851,
        provider = "EUNICE",
        month = 7,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.17131,
        provider = "NRG",
        month = 7,
        duration = 1
    ),
    Plan_t(
        name = "prime 4U",
        type = "YELLOW",
        price = 0.21109,
        provider = "NRG",
        month = 7,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.1881,
        provider = "OTE ESTATE",
        month = 7,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.1599,
        provider = "PROTERGIA",
        month = 7,
        duration = 1
    ),
    Plan_t(
        name = "Protergia Οικιακό - Value Easy",
        type = "YELLOW",
        price = 0.09063,
        provider = "PROTERGIA",
        month = 7,
        duration = 1
    ),
    Plan_t(
        name = "Protergia Οικιακό - Value Safe",
        type = "BLUE",
        price = 0.095,
        provider = "PROTERGIA",
        month = 7,
        duration = 6
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.17472,
        provider = "SOLAR ENERGY",
        month = 7,
        duration = 1
    ),
    Plan_t(
        name = "Οικιακό S1",
        type = "YELLOW",
        price = 0.20342,
        provider = "SOLAR ENERGY",
        month = 7,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.17213,
        provider = "VOLTERRA",
        month = 7,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.1588,
        provider = "VOLTON",
        month = 7,
        duration = 1
    ),
    Plan_t(
        name = "VOLTON YELLOW SIMPLE",
        type = "YELLOW",
        price = 0.1991,
        provider = "VOLTON",
        month = 7,
        duration = 1
    ),
    Plan_t(
        name = "VOLTON BLUE FIXED 6M_03",
        type = "BLUE",
        price = 0.144,
        provider = "VOLTON",
        month = 7,
        duration = 6
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.13557,
        provider = "ΔΕΗ",
        month = 7,
        duration = 1
    ),
    Plan_t(
        name = "MyHome 4All",
        type = "YELLOW",
        price = 0.12912,
        provider = "ΔΕΗ",
        month = 7,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.16637,
        provider = "ΕΛΙΝΟΙΛ",
        month = 7,
        duration = 1
    ),
    Plan_t(
        name = "Power On! Home Easy",
        type = "YELLOW",
        price = 0.19831,
        provider = "ΕΛΙΝΟΙΛ",
        month = 7,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.20525,
        provider = "ΖΕΝΙΘ",
        month = 7,
        duration = 1
    ),
    Plan_t(
        name = "Power Home Go Electric Plus",
        type = "YELLOW",
        price = 0.16033,
        provider = "ΖΕΝΙΘ",
        month = 7,
        duration = 1
    ),
    Plan_t(
        name = "Power Home Direct Plus",
        type = "BLUE",
        price = 0.095,
        provider = "ΖΕΝΙΘ",
        month = 7,
        duration = 6
    ),
    Plan_t(
        name = "BASIC HOME",
        type = "GREEN",
        price = 0.15879,
        provider = "ΗΡΩΝ",
        month = 7,
        duration = 1
    ),
    Plan_t(
        name = "YELLOW ONE HOME",
        type = "YELLOW",
        price = 0.0936,
        provider = "ΗΡΩΝ",
        month = 7,
        duration = 1
    ),
    Plan_t(
        name = "HOME FIX 5",
        type = "BLUE",
        price = 0.1336,
        provider = "ΗΡΩΝ",
        month = 7,
        duration = 6
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.1399,
        provider = "ΦΥΣΙΚΟ ΑΕΡΙΟ ΕΛΛΗΝΙΚΗ ΕΤΑΙΡΙΑ ΕΝΕΡΓΕΙΑΣ",
        month = 7,
        duration = 1
    ),
    Plan_t(
        name = "Maxi Home Super Save",
        type = "YELLOW",
        price = 0.077,
        provider = "ΦΥΣΙΚΟ ΑΕΡΙΟ ΕΛΛΗΝΙΚΗ ΕΤΑΙΡΙΑ ΕΝΕΡΓΕΙΑΣ",
        month = 7,
        duration = 1
    ),
    Plan_t(
        name = "Home Fixed",
        type = "BLUE",
        price = 0.199,
        provider = "ΦΥΣΙΚΟ ΑΕΡΙΟ ΕΛΛΗΝΙΚΗ ΕΤΑΙΡΙΑ ΕΝΕΡΓΕΙΑΣ",
        month = 7,
        duration = 3
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.1995,
        provider = "ELPEDISON",
        month = 8,
        duration = 1
    ),
    Plan_t(
        name = "ELPEDISON Win More",
        type = "YELLOW",
        price = 0.0895,
        provider = "ELPEDISON",
        month = 8,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.24365,
        provider = "EUNICE",
        month = 8,
        duration = 1
    ),
    Plan_t(
        name = "We Home",
        type = "YELLOW",
        price = 0.17911,
        provider = "EUNICE",
        month = 8,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.18731,
        provider = "NRG",
        month = 8,
        duration = 1
    ),
    Plan_t(
        name = "prime 4U",
        type = "YELLOW",
        price = 0.20124,
        provider = "NRG",
        month = 8,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.16811,
        provider = "OTE ESTATE",
        month = 8,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.1599,
        provider = "PROTERGIA",
        month = 8,
        duration = 1
    ),
    Plan_t(
        name = "Protergia Οικιακό - Value Easy",
        type = "YELLOW",
        price = 0.09093,
        provider = "PROTERGIA",
        month = 8,
        duration = 1
    ),
    Plan_t(
        name = "Protergia Οικιακό - Value Safe",
        type = "BLUE",
        price = 0.095,
        provider = "PROTERGIA",
        month = 8,
        duration = 6
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.19439,
        provider = "SOLAR ENERGY",
        month = 8,
        duration = 1
    ),
    Plan_t(
        name = "Οικιακό S1+",
        type = "YELLOW",
        price = 0.18038,
        provider = "SOLAR ENERGY",
        month = 8,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.17213,
        provider = "VOLTERRA",
        month = 8,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.17892,
        provider = "VOLTON",
        month = 8,
        duration = 1
    ),
    Plan_t(
        name = "VOLTON YELLOW SIMPLE",
        type = "YELLOW",
        price = 0.18919,
        provider = "VOLTON",
        month = 8,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.16391,
        provider = "ΔΕΗ",
        month = 8,
        duration = 1
    ),
    Plan_t(
        name = "MyHome 4All",
        type = "YELLOW",
        price = 0.15194,
        provider = "ΔΕΗ",
        month = 8,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.15805,
        provider = "ΕΛΙΝΟΙΛ",
        month = 8,
        duration = 1
    ),
    Plan_t(
        name = "Power On! Home Easy",
        type = "YELLOW",
        price = 0.17609,
        provider = "ΕΛΙΝΟΙΛ",
        month = 8,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.27918,
        provider = "ΖΕΝΙΘ",
        month = 8,
        duration = 1
    ),
    Plan_t(
        name = "Power Home Go Electric Plus",
        type = "YELLOW",
        price = 0.15095,
        provider = "ΖΕΝΙΘ",
        month = 8,
        duration = 1
    ),
    Plan_t(
        name = "BASIC HOME",
        type = "GREEN",
        price = 0.15879,
        provider = "ΗΡΩΝ",
        month = 8,
        duration = 1
    ),
    Plan_t(
        name = "YELLOW ONE HOME",
        type = "YELLOW",
        price = 0.0936,
        provider = "ΗΡΩΝ",
        month = 8,
        duration = 1
    ),
    Plan_t(
        name = "FIX GENIUS HOME 5",
        type = "BLUE",
        price = 0.1416,
        provider = "ΗΡΩΝ",
        month = 8,
        duration = 6
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.1699,
        provider = "ΦΥΣΙΚΟ ΑΕΡΙΟ ΕΛΛΗΝΙΚΗ ΕΤΑΙΡΙΑ ΕΝΕΡΓΕΙΑΣ",
        month = 8,
        duration = 1
    ),
    Plan_t(
        name = "Ρεύμα MAXI Home 10",
        type = "YELLOW",
        price = 0.18209,
        provider = "ΦΥΣΙΚΟ ΑΕΡΙΟ ΕΛΛΗΝΙΚΗ ΕΤΑΙΡΙΑ ΕΝΕΡΓΕΙΑΣ",
        month = 8,
        duration = 1
    ),
    Plan_t(
        name = "Home Fixed",
        type = "BLUE",
        price = 0.199,
        provider = "ΦΥΣΙΚΟ ΑΕΡΙΟ ΕΛΛΗΝΙΚΗ ΕΤΑΙΡΙΑ ΕΝΕΡΓΕΙΑΣ",
        month = 8,
        duration = 3
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.1995,
        provider = "ELPEDISON",
        month = 9,
        duration = 1
    ),
    Plan_t(
        name = "ELPEDISON Win More",
        type = "YELLOW",
        price = 0.0895,
        provider = "ELPEDISON",
        month = 9,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.19102,
        provider = "EUNICE",
        month = 9,
        duration = 1
    ),
    Plan_t(
        name = "We Home",
        type = "YELLOW",
        price = 0.15953,
        provider = "EUNICE",
        month = 9,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.18731,
        provider = "NRG",
        month = 9,
        duration = 1
    ),
    Plan_t(
        name = "nrg simple",
        type = "YELLOW",
        price = 0.089,
        provider = "NRG",
        month = 9,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.17424,
        provider = "OTE ESTATE",
        month = 9,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.1599,
        provider = "PROTERGIA",
        month = 9,
        duration = 1
    ),
    Plan_t(
        name = "Protergia Οικιακό - Value Easy",
        type = "YELLOW",
        price = 0.09093,
        provider = "PROTERGIA",
        month = 9,
        duration = 1
    ),
    Plan_t(
        name = "Protergia Οικιακό - Value Safe",
        type = "BLUE",
        price = 0.095,
        provider = "PROTERGIA",
        month = 9,
        duration = 6
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.16005,
        provider = "SOLAR ENERGY",
        month = 9,
        duration = 1
    ),
    Plan_t(
        name = "Οικιακό S1+",
        type = "YELLOW",
        price = 0.1596,
        provider = "SOLAR ENERGY",
        month = 9,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.17213,
        provider = "VOLTERRA",
        month = 9,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.19479,
        provider = "VOLTON",
        month = 9,
        duration = 1
    ),
    Plan_t(
        name = "VOLTON YELLOW SIMPLE",
        type = "YELLOW",
        price = 0.1746,
        provider = "VOLTON",
        month = 9,
        duration = 1
    ),
    Plan_t(
        name = "VOLTON BLUE FIXED 6M_04",
        type = "BLUE",
        price = 0.1592,
        provider = "VOLTON",
        month = 9,
        duration = 6
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.1594,
        provider = "ΔΕΗ",
        month = 9,
        duration = 1
    ),
    Plan_t(
        name = "MyHome 4All",
        type = "YELLOW",
        price = 0.14911,
        provider = "ΔΕΗ",
        month = 9,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.15987,
        provider = "ΕΛΙΝΟΙΛ",
        month = 9,
        duration = 1
    ),
    Plan_t(
        name = "Power On! Home Easy",
        type = "YELLOW",
        price = 0.15547,
        provider = "ΕΛΙΝΟΙΛ",
        month = 9,
        duration = 1
    ),
    Plan_t(
        name = "Power On! Blue Benefit",
        type = "BLUE",
        price = 0.125,
        provider = "ΕΛΙΝΟΙΛ",
        month = 9,
        duration = 3
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.15833,
        provider = "ΖΕΝΙΘ",
        month = 9,
        duration = 1
    ),
    Plan_t(
        name = "Power Home Go Electric Plus",
        type = "YELLOW",
        price = 0.13712,
        provider = "ΖΕΝΙΘ",
        month = 9,
        duration = 1
    ),
    Plan_t(
        name = "BASIC HOME",
        type = "GREEN",
        price = 0.15879,
        provider = "ΗΡΩΝ",
        month = 9,
        duration = 1
    ),
    Plan_t(
        name = "YELLOW ONE HOME",
        type = "YELLOW",
        price = 0.0936,
        provider = "ΗΡΩΝ",
        month = 9,
        duration = 1
    ),
    Plan_t(
        name = "BLUE GENEROUS HOME",
        type = "BLUE",
        price = 0.0936,
        provider = "ΗΡΩΝ",
        month = 9,
        duration = 6
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.1735,
        provider = "ΦΥΣΙΚΟ ΑΕΡΙΟ ΕΛΛΗΝΙΚΗ ΕΤΑΙΡΙΑ ΕΝΕΡΓΕΙΑΣ",
        month = 9,
        duration = 1
    ),
    Plan_t(
        name = "Ρεύμα MAXI Home Super Save+",
        type = "YELLOW",
        price = 0.072,
        provider = "ΦΥΣΙΚΟ ΑΕΡΙΟ ΕΛΛΗΝΙΚΗ ΕΤΑΙΡΙΑ ΕΝΕΡΓΕΙΑΣ",
        month = 9,
        duration = 1
    ),
    Plan_t(
        name = "Home Fixed",
        type = "BLUE",
        price = 0.199,
        provider = "ΦΥΣΙΚΟ ΑΕΡΙΟ ΕΛΛΗΝΙΚΗ ΕΤΑΙΡΙΑ ΕΝΕΡΓΕΙΑΣ",
        month = 9,
        duration = 3
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.18558,
        provider = "ELPEDISON",
        month = 10,
        duration = 1
    ),
    Plan_t(
        name = "ELPEDISON Win More 2",
        type = "YELLOW",
        price = 0.0949,
        provider = "ELPEDISON",
        month = 10,
        duration = 1
    ),
    Plan_t(
        name = "ELPEDISON Bright 6M",
        type = "BLUE",
        price = 0.0999,
        provider = "ELPEDISON",
        month = 10,
        duration = 6
    ),
    Plan_t(
        name = "ELPEDISON Bright 12M",
        type = "BLUE",
        price = 0.1049,
        provider = "ELPEDISON",
        month = 10,
        duration = 12
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.15785,
        provider = "EUNICE",
        month = 10,
        duration = 1
    ),
    Plan_t(
        name = "Eunice Βασικός Τιμοκατάλογος",
        type = "YELLOW",
        price = 0.13456,
        provider = "EUNICE",
        month = 10,
        duration = 1
    ),
    Plan_t(
        name = "Eunice Home Fair",
        type = "BLUE",
        price = 0.145,
        provider = "EUNICE",
        month = 10,
        duration = 3
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.15838,
        provider = "NRG",
        month = 10,
        duration = 1
    ),
    Plan_t(
        name = "nrg simple",
        type = "YELLOW",
        price = 0.089,
        provider = "NRG",
        month = 10,
        duration = 1
    ),
    Plan_t(
        name = "nrg fixed 4U 12μήνες",
        type = "BLUE",
        price = 0.189,
        provider = "NRG",
        month = 10,
        duration = 12
    ),
    Plan_t(
        name = "nrg fixed 4U 6μήνες",
        type = "BLUE",
        price = 0.095,
        provider = "NRG",
        month = 10,
        duration = 6
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.1683,
        provider = "OTE ESTATE",
        month = 10,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.1599,
        provider = "PROTERGIA",
        month = 10,
        duration = 1
    ),
    Plan_t(
        name = "Protergia Οικιακό - Value Easy",
        type = "YELLOW",
        price = 0.09093,
        provider = "PROTERGIA",
        month = 10,
        duration = 1
    ),
    Plan_t(
        name = "Value Safe More",
        type = "BLUE",
        price = 0.125,
        provider = "PROTERGIA",
        month = 10,
        duration = 6
    ),
    Plan_t(
        name = "Οικιακό S1+",
        type = "YELLOW",
        price = 0.1331,
        provider = "SOLAR ENERGY",
        month = 10,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.14484,
        provider = "VOLTERRA",
        month = 10,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.17894,
        provider = "VOLTON",
        month = 10,
        duration = 1
    ),
    Plan_t(
        name = "VOLTON YELLOW SIMPLE",
        type = "YELLOW",
        price = 0.15051,
        provider = "VOLTON",
        month = 10,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.1496,
        provider = "ΔΕΗ",
        month = 10,
        duration = 1
    ),
    Plan_t(
        name = "MyHome 4All",
        type = "YELLOW",
        price = 0.13852,
        provider = "ΔΕΗ",
        month = 10,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.13974,
        provider = "ΕΛΙΝΟΙΛ",
        month = 10,
        duration = 1
    ),
    Plan_t(
        name = "Power On! Home Easy",
        type = "YELLOW",
        price = 0.12921,
        provider = "ΕΛΙΝΟΙΛ",
        month = 10,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.14904,
        provider = "ΖΕΝΙΘ",
        month = 10,
        duration = 1
    ),
    Plan_t(
        name = "BASIC HOME",
        type = "GREEN",
        price = 0.15879,
        provider = "ΗΡΩΝ",
        month = 10,
        duration = 1
    ),
    Plan_t(
        name = "YELLOW ONE HOME",
        type = "YELLOW",
        price = 0.13754,
        provider = "ΗΡΩΝ",
        month = 10,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.15915,
        provider = "ΦΥΣΙΚΟ ΑΕΡΙΟ ΕΛΛΗΝΙΚΗ ΕΤΑΙΡΙΑ ΕΝΕΡΓΕΙΑΣ",
        month = 10,
        duration = 1
    ),
    Plan_t(
        name = "Ρεύμα MAXI Home Super Save+ Oct",
        type = "YELLOW",
        price = 0.072,
        provider = "ΦΥΣΙΚΟ ΑΕΡΙΟ ΕΛΛΗΝΙΚΗ ΕΤΑΙΡΙΑ ΕΝΕΡΓΕΙΑΣ",
        month = 10,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.1487,
        provider = "ELPEDISON",
        month = 11,
        duration = 1
    ),
    Plan_t(
        name = "ELPEDISON Win More 2",
        type = "YELLOW",
        price = 0.0949,
        provider = "ELPEDISON",
        month = 11,
        duration = 1
    ),
    Plan_t(
        name = "ELPEDISON Bright 6M",
        type = "BLUE",
        price = 0.0999,
        provider = "ELPEDISON",
        month = 11,
        duration = 6
    ),
    Plan_t(
        name = "ELPEDISON Bright 12M",
        type = "BLUE",
        price = 0.1049,
        provider = "ELPEDISON",
        month = 11,
        duration = 12
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.12748,
        provider = "EUNICE",
        month = 11,
        duration = 1
    ),
    Plan_t(
        name = "Eunice Home Flow",
        type = "YELLOW",
        price = 0.11223,
        provider = "EUNICE",
        month = 11,
        duration = 1
    ),
    Plan_t(
        name = "Eunice Home Fair",
        type = "BLUE",
        price = 0.145,
        provider = "EUNICE",
        month = 11,
        duration = 3
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.1486,
        provider = "NRG",
        month = 11,
        duration = 1
    ),
    Plan_t(
        name = "nrg simple 1.0",
        type = "YELLOW",
        price = 0.089,
        provider = "NRG",
        month = 11,
        duration = 1
    ),
    Plan_t(
        name = "nrg fixed 4U 12μήνες",
        type = "BLUE",
        price = 0.189,
        provider = "NRG",
        month = 11,
        duration = 12
    ),
    Plan_t(
        name = "nrg fixed 4U 6μήνες",
        type = "BLUE",
        price = 0.095,
        provider = "NRG",
        month = 11,
        duration = 6
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.16236,
        provider = "OTE ESTATE",
        month = 11,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.149,
        provider = "PROTERGIA",
        month = 11,
        duration = 1
    ),
    Plan_t(
        name = "Protergia Οικιακό - Value Easy",
        type = "YELLOW",
        price = 0.09093,
        provider = "PROTERGIA",
        month = 11,
        duration = 1
    ),
    Plan_t(
        name = "Value Safe More",
        type = "BLUE",
        price = 0.125,
        provider = "PROTERGIA",
        month = 11,
        duration = 6
    ),
    Plan_t(
        name = "Protergia Οικιακό - Value Safe",
        type = "BLUE",
        price = 0.095,
        provider = "PROTERGIA",
        month = 11,
        duration = 6
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.13794,
        provider = "SOLAR ENERGY",
        month = 11,
        duration = 1
    ),
    Plan_t(
        name = "Οικιακό FIXED",
        type = "BLUE",
        price = 0.36,
        provider = "SOLAR ENERGY",
        month = 11,
        duration = 3
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.145,
        provider = "VOLTERRA",
        month = 11,
        duration = 1
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.15185,
        provider = "VOLTON",
        month = 11,
        duration = 1
    ),
    Plan_t(
        name = "VOLTON YELLOW SIMPLE",
        type = "YELLOW",
        price = 0.12409,
        provider = "VOLTON",
        month = 11,
        duration = 1
    ),
    Plan_t(
        name = "VOLTON BLUE FIXED 6M_04",
        type = "BLUE",
        price = 0.1592,
        provider = "VOLTON",
        month = 11,
        duration = 6
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.1408,
        provider = "ΔΕΗ",
        month = 11,
        duration = 1
    ),
    Plan_t(
        name = "MyHome 4All",
        type = "YELLOW",
        price = 0.13367,
        provider = "ΔΕΗ",
        month = 11,
        duration = 1
    ),
    Plan_t(
        name = "My Home Online",
        type = "BLUE",
        price = 0.142,
        provider = "ΔΕΗ",
        month = 11,
        duration = 6
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.11666,
        provider = "ΕΛΙΝΟΙΛ",
        month = 11,
        duration = 1
    ),
    Plan_t(
        name = "Power On! Blue",
        type = "BLUE",
        price = 0.149,
        provider = "ΕΛΙΝΟΙΛ",
        month = 11,
        duration = 3
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.13682,
        provider = "ΖΕΝΙΘ",
        month = 11,
        duration = 1
    ),
    Plan_t(
        name = "Power Home Go Electric Plus",
        type = "YELLOW",
        price = 0.1143,
        provider = "ΖΕΝΙΘ",
        month = 11,
        duration = 1
    ),
    Plan_t(
        name = "Power Home Direct",
        type = "BLUE",
        price = 0.129,
        provider = "ΖΕΝΙΘ",
        month = 11,
        duration = 6
    ),
    Plan_t(
        name = "BASIC HOME",
        type = "GREEN",
        price = 0.13737,
        provider = "ΗΡΩΝ",
        month = 11,
        duration = 1
    ),
    Plan_t(
        name = "YELLOW ONE HOME",
        type = "YELLOW",
        price = 0.10337,
        provider = "ΗΡΩΝ",
        month = 11,
        duration = 1
    ),
    Plan_t(
        name = "BLUE GENEROUS HOME",
        type = "BLUE",
        price = 0.0936,
        provider = "ΗΡΩΝ",
        month = 11,
        duration = 6
    ),
    Plan_t(
        name = "Ειδικό",
        type = "GREEN",
        price = 0.14681,
        provider = "ΦΥΣΙΚΟ ΑΕΡΙΟ ΕΛΛΗΝΙΚΗ ΕΤΑΙΡΙΑ ΕΝΕΡΓΕΙΑΣ",
        month = 11,
        duration = 1
    ),
    Plan_t(
        name = "Ρεύμα MAXI Home Economy XL",
        type = "YELLOW",
        price = 0.13503,
        provider = "ΦΥΣΙΚΟ ΑΕΡΙΟ ΕΛΛΗΝΙΚΗ ΕΤΑΙΡΙΑ ΕΝΕΡΓΕΙΑΣ",
        month = 11,
        duration = 1
    ),
    Plan_t(
        name = "Ρεύμα Blue Friday",
        type = "BLUE",
        price = 0.089,
        provider = "ΦΥΣΙΚΟ ΑΕΡΙΟ ΕΛΛΗΝΙΚΗ ΕΤΑΙΡΙΑ ΕΝΕΡΓΕΙΑΣ",
        month = 11,
        duration = 6
    )
]
