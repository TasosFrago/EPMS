import logging
import os

import pymysql
from sshtunnel import SSHTunnelForwarder

from consumer_data import  getConsumerData
from meter_data import getMeterData, gen_meterNum

from typing import Final
from dataclasses import dataclass
from contextlib import contextmanager
from dotenv import load_dotenv

# logging.basicConfig(level=logging.DEBUG)

CONSUMER_NUMBER: Final[int] = 2

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
        self.db_name = "lab2425omada1_EPMS"
        self.db_port = 3306

config = Config()


@contextmanager
def sshtunnelAndMySQLconn(config):
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
            port=localport,
            charset="utf8mb4"
        )
        try:
            yield connection
        finally:
            connection.commit()
            print("Closing DB connection.")
            connection.close()
            print("Closed DB connection")
        print("SSH connection closed.")

def loadTBL_CONSUMER_METER(connection):
    curs = connection.cursor()
    for _ in range(CONSUMER_NUMBER):
        cons = getConsumerData()
        landline = [', landline', f', "{cons.landline}"'] if cons.landline else ["", ""]

        curs.execute(f"""
        INSERT INTO CONSUMER
        (first_name, last_name, email, cell{landline[0]})
        VALUES
        ("{cons.first_name}", "{cons.last_name}", "{cons.email}", "{cons.cell}" {landline[1]});"""
        )
        curs.execute("SELECT LAST_INSERT_ID();")

        ## Load METER table
        for _ in range(gen_meterNum()):
            meter = getMeterData(int(curs.fetchone()[0])) # type: ignore I hate people

            curs.execute(f"""
            INSERT INTO METER
            (status, kWh, address, rated_power, owner)
            VALUES
            ("{meter.status}", "{meter.kWh}", "{meter.address}", "{meter.rated_power}", "{meter.owner}")
            """)


with sshtunnelAndMySQLconn(config) as conn:
    cursorclass = pymysql.cursors.DictCursor # type: ignore Fucking amatures can't fucking type hint their shitty little libray
    curs = conn.cursor()
    curs.execute("SHOW TABLES;")
    for table in curs.fetchall():
        print(f"{table=}")

    ## SAFEGUARD delete later
    curs.execute("DELETE FROM METER;") # First delete METER because it has foreign keys
    curs.execute("DELETE FROM CONSUMER;")

    loadTBL_CONSUMER_METER(conn)

    curs.execute("SELECT * FROM CONSUMER")
    for consumer in curs.fetchall():
        print(consumer)

    curs.execute("SELECT * FROM METER")
    for consumer in curs.fetchall():
        print(consumer)
