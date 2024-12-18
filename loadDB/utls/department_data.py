from pydantic import BaseModel, field_validator

class Department_t(BaseModel):
    region: str
    phone: str # 2-digit area code, 8-digit random
    
    @field_validator("phone")
    @classmethod
    def validate_phone(cls, v: str):
        valid_prefixes = ["21", "22", "23", "24", "25", "26", "27", "28"]
        if len(v) != 10 or not v.isdigit():
            raise ValueError(f"Invalid phone number: {v}. Must be 10 digits.")
        if not v.startswith(valid_prefixes):
            raise ValueError(f"Invalid phone number: {v}. Must start with one of the following prefixes: {', '.join(valid_prefixes)}.")
        return v

departments: list[Department_t] = [
    Department_t(
        region = "Αθήνα",
        phone = "2102525256"
    ),
    Department_t(
        region = "Ανατολική Αττική",
        phone = "2156567890"
    ),
    Department_t(
        region = "Δυτική Αττική",
        phone = "2122334455"
    ),
    Department_t(
        region = "Πειραιάς",
        phone = "2123234567"
    ),
    Department_t(
        region = "Εύβοια",
        phone = "2245454567"
    ),
    Department_t(
        region = "Ευρυτανία",
        phone = "2256567812"
    ),
    Department_t(
        region = "Φωκίδα",
        phone = "2234344556"
    ),
    Department_t(
        region = "Φθιώτιδα",
        phone = "2223552355"
    ),
    Department_t(
        region = "Βοιωτία",
        phone = "2257571244"
    ),
    Department_t(
        region = "Χαλκιδική",
        phone = "2350070080"
    ),
    Department_t(
        region = "Ημαθία",
        phone = "2390905711"
    ),
    Department_t(
        region = "Κιλκίς",
        phone = "2313700800"
    ),
    Department_t(
        region = "Πέλλα",
        phone = "2340504070"
    ),
    Department_t(
        region = "Πιερία",
        phone = "2344552233"
    ),
    Department_t(
        region = "Σέρρες",
        phone = "2341512000"
    ),
    Department_t(
        region = "Θεσσαλονίκη",
        phone = "2310600100"
    ),
    Department_t(
        region = "Χανιά",
        phone = "2817174455"
    ),
    Department_t(
        region = "Ηράκλειο",
        phone = "2822335050"
    ),
    Department_t(
        region = "Λασίθι",
        phone = "2835356060"
    ),
    Department_t(
        region = "Ρέθυμνο",
        phone = "2870060055"
    ),
    Department_t(
        region = "Δράμα",
        phone = "2530201000"
    ),
    Department_t(
        region = "Έβρος",
        phone = "2566774242"
    ),
    Department_t(
        region = "Καβάλα",
        phone = "2534503450"
    ),
    Department_t(
        region = "Ροδόπη",
        phone = "2510050700"
    ),
    Department_t(
        region = "Ξάνθη",
        phone = "2541500800"
    ),
    Department_t(
        region = "Άρτα",
        phone = "2632451234"
    ),
    Department_t(
        region = "Ιωάννινα",
        phone = "2640405050"
    ),
    Department_t(
        region = "Πρέβεζα",
        phone = "2657843230"
    ),
    Department_t(
        region = "Θεσπρωτία",
        phone = "2613572468"
    ),
    Department_t(
        region = "Κέρκυρα",
        phone = "2686427531"
    ),
    Department_t(
        region = "Κεφαλονιά",
        phone = "2618467091"
    ),
    Department_t(
        region = "Λευκάδα",
        phone = "2633447788"
    ),
    Department_t(
        region = "Ζάκυνθος",
        phone = "2650600700"
    ),
    Department_t(
        region = "Χίος",
        phone = "2260400300"
    ),
    Department_t(
        region = "Λέσβος",
        phone = "2205060304"
    ),
    Department_t(
        region = "Σάμος",
        phone = "2250098764"
    ),
    Department_t(
        region = "Αρκαδία",
        phone = "2710020060"
    ),
    Department_t(
        region = "Αργολίδα",
        phone = "2755577722"
    ),
    Department_t(
        region = "Κόρινθος",
        phone = "2711160090"
    ),
    Department_t(
        region = "Λακωνία",
        phone = "2755533444"
    ),
    Department_t(
        region = "Μεσσήνη",
        phone = "2799977788"
    ),
    Department_t(
        region = "Κυκλάδες",
        phone = "2244993333"
    ),
    Department_t(
        region = "Δωδεκάνησα",
        phone = "2233366222"
    ),
    Department_t(
        region = "Καρδίτσα",
        phone = "2477733555"
    ),
    Department_t(
        region = "Λάρισα",
        phone = "2422266111"
    ),
    Department_t(
        region = "Μαγνησία",
        phone = "2466999555"
    ),
    Department_t(
        region = "Τρίκαλα",
        phone = "2466444333"
    ),
    Department_t(
        region = "Αχαΐα",
        phone = "2622888111"
    ),
    Department_t(
        region = "Αιτωλοακαρνανία",
        phone = "2699000111"
    ),
    Department_t(
        region = "Ηλεία",
        phone = "2677000222"
    ),
    Department_t(
        region = "Φλώρινα",
        phone = "2300444665"
    ),
    Department_t(
        region = "Γρεβενά",
        phone = "2467823451"
    ),
    Department_t(
        region = "Καστοριά",
        phone = "24"
    ),
    Department_t(
        region = "Κοζάνη",
        phone = "24"
    ),
    Department_t(
        region = "Άγιο Όρος",
        phone = "2322420666"
    )
]
