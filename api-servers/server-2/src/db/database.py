import psycopg2

from src.api.settings import settings
from src.api.exceptions import ApiFailedConnectDataBase
from pydantic import BaseModel
from typing import Union


class PostgresConnection():
    def __init__(self):
        self.user = settings.DB_USER
        self.password = settings.DB_PASSWORD
        self.port = settings.DB_PORT
        self.database = settings.DB_DATABSE

        self.conn: Union[None, psycopg2.connect]

    def _connect_db(self):
        try:
            self.conn = psycopg2.connect(
                user=self.user,
                password=self.password,
                host="127.0.0.1",
                port=self.port,
                database=self.database
            )
        except Exception as e:
            print(e)
            raise ApiFailedConnectDataBase

    def _close_connection(self) -> None:
        if self.conn is not None:
            self.conn.close()
