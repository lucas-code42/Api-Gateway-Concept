from src.db.db import Postgres
from psycopg2 import connection


class BookRepository:
    def __init__(self, db_connection: connection) -> None:
        self.pg_connection = db_connection
        self.cursor = self.pg_connection.cursor()

    def create_book(self):
        self.cursor.execute("")
