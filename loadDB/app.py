import logging
import os

import pymysql
from sshtunnel import SSHTunnelForwarder

from dataclasses import dataclass
from contextlib import contextmanager
from dotenv import load_dotenv

# logging.basicConfig(level=logging.DEBUG)

load_dotenv()

@dataclass
class Config:
    """SSH connection configurations"""
    usrname: str
    passwd: str
    server_host: str
    server_port: int
    db_host: str
    db_name: str
    db_port: int

    def __init__(self):
        self.usrname = str(os.getenv("USERNAME"))
        self.passwd = str(os.getenv("PASSWORD"))
        self.server_host = str(os.getenv("HOST"))
        self.server_port = int(os.getenv("PORT", 22210))
        self.db_host = "127.0.0.1"
        self.db_name = "lab2425omada1_Company"
        self.db_port = 3306

config = Config()


@contextmanager
def sshtunnelAndDBconnection(config):
    with SSHTunnelForwarder(
        (config.server_host, config.server_port),
        ssh_username=config.usrname,
        ssh_password=config.passwd,
        remote_bind_address=(config.db_host, config.db_port)
    ) as tunnel:
        print("Established ssh connection.")
        localport: int = tunnel.local_bind_port # type: ignore

        connection = pymysql.connect(
            host=config.db_host,
            user=config.usrname,
            passwd=config.passwd,
            db=config.db_name,
            port=localport
        )
        try:
            yield connection
        finally:
            print("Closing DB connection.")
            connection.close()
            print("Closed DB connection")
        print("SSH connection closed.")

with sshtunnelAndDBconnection(config) as conn:
    curs = conn.cursor()
    curs.execute("SHOW TABLES;")
    for table in curs.fetchall():
        print(f"{table=}")
