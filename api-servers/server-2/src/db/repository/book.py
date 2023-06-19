from src.api.model.books import BooksModels
from src.api.exceptions import ApiFailedToInsertBook

import psycopg2


class BookRepository:
    def __init__(self, db_connection: psycopg2.connect) -> None:
        self.pg_connection = db_connection
        self.cursor = self.pg_connection.cursor()

    def create_book(self, book: BooksModels) -> bool:
        result = False
        try:
            query = """
                INSERT INTO 
                    books (name, price, author)
                VALUES 
                    (%s, %s,%s)
            """
            insert_data = (book.name, book.price, book.author)

            self.cursor.execute(query, insert_data)
            self.pg_connection.commit()

            if self.cursor.rowcount > 0:
                result = True
        except Exception:
            raise ApiFailedToInsertBook
        finally:
            if self.cursor:
                self.cursor.close()

        return result
