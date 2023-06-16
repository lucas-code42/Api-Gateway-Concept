import psycopg2
from psycopg2 import connection
from src.api.settings import settings


class Postgres:
    def __init__(self):
        self.user = settings.DB_USER
        self.password = settings.DB_PASSWORD
        self.port = settings.DB_PORT
        self.database = settings.DB_DATABSE

    def connect_db(self) -> (connection | None):
        conn = None
        try:
            conn = psycopg2.connect(
                user=self.user,
                password=self.password,
                host="127.0.0.1",
                port=self.port,
                database=self.database
            )
            return conn
        except:
            return conn
