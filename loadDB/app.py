import logging
import os

import random
import pymysql
from sshtunnel import SSHTunnelForwarder

from consumer_data import  getConsumerData
from employee_data import getEmployeeData
from meter_data import getMeterData, gen_meterNum, gen_kWh
from provider_data import providers
from plan_data import plans
from month_data import months

from typing import Final
from dataclasses import dataclass, field
from contextlib import contextmanager
from dotenv import load_dotenv
from pathlib import Path

# logging.basicConfig(level=logging.DEBUG)

CONSUMER_NUMBER: Final[int] = 100
EMPLOYEE_NUMBER: Final[int] = 5

Connection_t = pymysql.connections.Connection

dotenv_path = Path("../.env")

load_dotenv(dotenv_path=dotenv_path, override=True)
@dataclass(frozen=True)
class TerminalColors:
    # Reset
    RESET: str = "\033[0m"
    # Text Colors
    BLACK: str = "\033[30m"
    RED: str = "\033[31m"
    GREEN: str = "\033[32m"
    YELLOW: str = "\033[33m"
    BLUE: str = "\033[34m"
    MAGENTA: str = "\033[35m"
    CYAN: str = "\033[36m"
    WHITE: str = "\033[37m"
termC = TerminalColors()

@dataclass
class Config:
    """SSH connection configurations"""
    usrname: str = field(default_factory=lambda: str(os.getenv("USERNAME")))
    passwd: str = field(default_factory=lambda: str(os.getenv("PASSWORD")))
    server_host: str = field(default_factory=lambda: str(os.getenv("HOST")))
    server_port: int = field(default_factory=lambda: int(os.getenv("PORT", 22210)))
    db_host: str = "127.0.0.1"
    db_name: str = "lab2425omada1_EPMS"
    db_port: int = 3306

config = Config()

@contextmanager
def sshtunnelAndMySQLconn(config):
    with SSHTunnelForwarder(
        (config.server_host, config.server_port),
        ssh_username=config.usrname,
        ssh_password=config.passwd,
        remote_bind_address=(config.db_host, config.db_port)
    ) as tunnel:
        print(f"{termC.GREEN}Established ssh connection.{termC.RESET}")
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
            print(f"{termC.YELLOW}Closing DB connection.{termC.RESET}")
            connection.close()
            print(f"{termC.GREEN}Closed DB connection.{termC.RESET}")
        print(f"{termC.WHITE}SSH connection closed.{termC.RESET}")

def loadTBL_CONSUMER_METER(connection: Connection_t) -> None:
    print("Starting loading tables CONSUMER and METER...")
    curs = connection.cursor()
    for _ in range(CONSUMER_NUMBER):
        cons = getConsumerData()
        landline = [', landline', f', "{cons.landline}"'] if cons.landline else ["", ""]

        curs.execute(f"""
        INSERT INTO CONSUMER
        (first_name, last_name, email, password, cell{landline[0]})
        VALUES
        ("{cons.first_name}", "{cons.last_name}", "{cons.email}", "{cons.password}", "{cons.cell}" {landline[1]});
        """)
        curs.execute("SELECT LAST_INSERT_ID();")
        owner_id = int(curs.fetchone()[0])

        ## Load METER table
        for _ in range(gen_meterNum()):
            meter = getMeterData(owner_id)

            curs.execute(f"""
            INSERT INTO METER
            (address, rated_power, owner)
            VALUES
            ("{meter.address}", {meter.rated_power}, {meter.owner});
            """)
    print(f"{termC.GREEN}Loaded tables CONSUMER and METER!{termC.RESET}")
    return

def loadTBL_EMPLOYEE(connection: Connection_t) -> None:
    print("Starting loading table EMPLOYEE...")
    curs = connection.cursor()

    for _ in range(EMPLOYEE_NUMBER):
        employee = getEmployeeData()
        salary = [', salary', f',"{employee.salary}"'] if employee.salary else ["",""]
        curs.execute(f"""
        INSERT INTO EMPLOYEE
        (first_name, last_name, email, password, phone {salary[0]})
        VALUES
        ("{employee.first_name}", "{employee.last_name}", "{employee.email}", "{employee.password}", "{employee.phone}" {salary[1]})
         """)
    print(f"{termC.GREEN}Loaded table EMPLOYEE!{termC.RESET}")
    return

def loadTBL_PROVIDER(connection: Connection_t) -> None:
    print("Starting loading table PROVIDER...")
    curs = connection.cursor()
    for provider in providers:
        curs.execute(f"""
        INSERT INTO PROVIDER
        (name, phone, email)
        VALUES
        ("{provider.name}", "{provider.phone}", "{provider.email}");
        """)
    print(f"{termC.GREEN}Loaded table PROVIDER!{termC.RESET}")
    return

def loadTBL_PLAN(connection: Connection_t) -> None:
    print("Starting loading table PLAN...")
    curs = connection.cursor()
    for plan in plans:
        curs.execute(f"""
        INSERT INTO PLAN
        (type, price, name, provider, month, year, duration)
        VALUES
        ("{plan.type}", {plan.price}, "{plan.name}", "{plan.provider}", "{months[plan.month-1].name}", {months[plan.month-1].year}, {plan.duration});
        """)
    print(f"{termC.GREEN}Loaded table PLAN!{termC.RESET}")
    return

def loadTBL_CHOOSES_INVOICE_PAYS(connection: Connection_t) -> None:
    print("Starting loading tables CHOOSES, INVOICE and PAYS...")
    cursorclass = pymysql.cursors.DictCursor # type: ignore My bad it was pyrights fault
    curs = connection.cursor(cursorclass)
    curs.execute(f"""
    SELECT supply_id, owner
    FROM METER;
    """)
    meter_info = curs.fetchall()
    for i in range(0, 10):
        curs.execute(f"""
        SELECT plan_id, provider, price
        FROM PLAN
        WHERE year = {months[i].year} AND month = "{months[i].name}";
        """)
        available_plans = curs.fetchall()
        for meter in meter_info:
            random_plan = available_plans[random.randrange(0, len(available_plans))]
            if i == 0:
                curs.execute(f"""
                INSERT INTO CHOOSES
                (user, plan, supply_id)
                VALUES
                ({meter["owner"]}, {random_plan["plan_id"]}, {meter["supply_id"]});
                """)
                curs.execute(f"""
                UPDATE METER
                SET status = {int(True)}
                WHERE supply_id = {meter["supply_id"]};
                """)
            else:
                curs.execute(f"""
                SELECT plan_id, month, year, duration
                FROM CHOOSES, PLAN
                WHERE plan = plan_id AND supply_id = {meter["supply_id"]};
                """)
                prev_plan = curs.fetchone()
                pos = [ind for ind, month in enumerate(months) if (month.name == prev_plan["month"]) and (month.year == prev_plan["year"])][0]

                if (i < (pos + prev_plan["duration"])):
                    random_plan["plan_id"] = prev_plan["plan_id"]

                curs.execute(f"""
                UPDATE CHOOSES
                SET plan = {random_plan["plan_id"]}
                WHERE supply_id = {meter["supply_id"]};
                """)
            kWh = gen_kWh()
            cost: float = kWh * random_plan["price"]
            curs.execute(f"""
            UPDATE METER
            SET kWh = {kWh}
            WHERE supply_id = {meter["supply_id"]};
            """)
            curs.execute(f"""
            INSERT INTO INVOICE
            (total, current_cost, receiver, provider, meter, plan)
            VALUES
            ({cost}, {cost}, {meter["owner"]}, "{random_plan["provider"]}", {meter["supply_id"]}, {random_plan["plan_id"]});
            """)
            curs.execute(f"""
            INSERT INTO PAYS
            (user, provider, supply_id, amount)
            VALUES
            ({meter["owner"]}, "{random_plan["provider"]}", {meter["supply_id"]}, {cost});
            """)
    print(f"{termC.GREEN}Loaded tables CHOOSES, INVOICE and PAYS!{termC.RESET}")
    return

def main():
    with sshtunnelAndMySQLconn(config) as conn:
        curs = conn.cursor()
        curs.execute("SHOW TABLES;")
        for table in curs.fetchall():
            print(f"{table=}")

        ## DELETES 
        ##curs.execute("DELETE FROM INVOICE;")
        ##curs.execute("DELETE FROM CHOOSES;")
        ##curs.execute("DELETE FROM PAYS;")
        curs.execute("DELETE FROM METER;") # First delete METER because it has foreign keys
        curs.execute("DELETE FROM CONSUMER;")
        #curs.execute("DELETE FROM PLAN;")
        #curs.execute("DELETE FROM PROVIDER;")

        ## INSERT DATA
        loadTBL_CONSUMER_METER(conn)
        #loadTBL_PROVIDER(conn)
        #loadTBL_PLAN(conn)
        #loadTBL_CHOOSES_INVOICE_PAYS(conn)

        printTBL = lambda tbl: [print(f"{termC.YELLOW}{row}{termC.RESET}") for row in (curs.execute(f"SELECT * FROM {tbl}"), curs.fetchall())[1]]

        printTBL("CONSUMER")
        printTBL("METER")
        # printTBL("PROVIDER")
        # printTBL("PLAN")
        # printTBL("INVOICE")

if __name__=="__main__":
    main()
