import string
from random import choice
from typing import Final

from passlib.context import CryptContext

PASSWORD_FILE: Final[str] = "passwd.txt"

PWD_CONTEXT=CryptContext(schemes=["bcrypt"], deprecated="auto")

def get_password(email: str):
    letters = string.ascii_lowercase
    passwd = "".join(choice(letters) for _ in range(10)) 

    with open(PASSWORD_FILE, "a") as fh:
        fh.write(f"{passwd} -- {email}\n")
    return PWD_CONTEXT.hash(passwd)
