from fastapi import APIRouter
from src.api.exceptions import ApiFailedConnectDataBase
from src.api.model.books import BooksModels
from src.db.db import PostgresConnection
from src.db.repository.book import BookRepository

create = APIRouter()


@create.post("/", response_model=BooksModels, status_code=200)
async def create_handler(book: BooksModels):
    pg = None
    response = None

    try:
        pg = PostgresConnection()
        pg._connect_db()

        repository = BookRepository(pg.conn)
        if repository.create_book(book):
            response = book
        else:
            response = {}

    except ApiFailedConnectDataBase:
        return "erro"
    finally:
        if isinstance(pg, PostgresConnection):
            pg._close_connection()

    return response
