from src.api.model.books import BooksModels
from src.api.exceptions import (
    ApiFailedToInsertBook, ApiFailedToDeleteBook, ApiFailedToGetBookById, ApiFailedToUpdateBooks, ApiFailedToVerifyId)
import psycopg2
from typing import List, Union


class BookRepository:
    def __init__(self, db_connection: psycopg2.connect) -> None:
        self.pg_connection: psycopg2.connection = db_connection
        self.cursor = self.pg_connection.cursor()

    async def create_book(self, book: BooksModels) -> int:
        result = None
        try:
            query = """
                INSERT INTO 
                    books (name, price, author)
                VALUES 
                    (%s, %s, %s)
                RETURNING id
            """
            insert_data = (book.name, book.price, book.author)
            self.cursor.execute(query, insert_data)
            self.pg_connection.commit()
            if self.cursor.rowcount > 0:
                result = self.cursor.fetchone()[0]
        except Exception:
            raise ApiFailedToInsertBook
        finally:
            if self.cursor:
                self.cursor.close()

        return result

    async def delete_book_by_id(self, book_id: int) -> bool:
        result = False
        try:
            query = """
                DELETE FROM
                    books
                WHERE
                    id = %s
            """
            self.cursor.execute(query, (book_id,))
            self.pg_connection.commit()

            print(self.cursor.rowcount)

            if self.cursor.rowcount > 0:
                result = True
        except Exception:
            raise ApiFailedToDeleteBook
        finally:
            if self.cursor:
                self.cursor.close()
        return result

    async def get_book_by_id(self, book_id: int) -> List[BooksModels]:
        result = None
        try:
            query = """
                SELECT * FROM 
                    books 
                WHERE 
                    books.id = %s;
            """
            self.cursor.execute(query, (book_id,))
            self.pg_connection.commit()
            result_buffer = self.cursor.fetchone()
            result = BooksModels(
                book_id=result_buffer[0],
                name=result_buffer[1],
                price=result_buffer[2],
                author=result_buffer[-1]
            )
        except Exception:
            raise ApiFailedToGetBookById
        finally:
            if self.cursor:
                self.cursor.close()
        return result

    async def get_all_books(self) -> List[BooksModels]:
        result = []
        try:
            query = """
                SELECT * FROM books
            """
            self.cursor.execute(query)
            self.pg_connection.commit()
            result_buffer = self.cursor.fetchall()
            for i in result_buffer:
                result.append(
                    BooksModels(
                        book_id=i[0],
                        name=i[1],
                        price=i[2],
                        author=i[3]
                    )
                )
        except Exception:
            raise ApiFailedToGetBookById
        finally:
            if self.cursor:
                self.cursor.close()
        return result

    async def verify_id(self, book_id: int) -> Union[int, None]:
        result = None
        try:
            query = """
                SELECT id FROM 
                    books
                WHERE 
                    id = %s;
            """
            self.cursor.execute(query, (str(book_id),))
            self.pg_connection.commit()
            buffer_id = self.cursor.fetchone()[0]
            if buffer_id > 0:
                result = buffer_id
        except Exception:
            raise ApiFailedToVerifyId
        return result

    async def udpate_book_by_id(self, book: BooksModels) -> BooksModels:
        try:
            verified_id = await self.verify_id(book.book_id)
            if verified_id is not None and book.book_id == verified_id:
                query = """
                    UPDATE books
                    SET
                        name = %s,
                        author = %s,
                        price = %s
                    WHERE 
                        id = %s;
                """
                self.cursor.execute(
                    query, (book.name, book.author, book.price, book.book_id))
                self.pg_connection.commit()
                if self.cursor.rowcount <= 0:
                    raise ApiFailedToUpdateBooks
            else:
                raise ApiFailedToVerifyId
        except Exception:
            raise ApiFailedToUpdateBooks
        finally:
            if self.cursor:
                self.cursor.close()
        return book
