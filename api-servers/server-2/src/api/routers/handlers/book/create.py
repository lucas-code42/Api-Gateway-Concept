from fastapi import APIRouter, HTTPException
from fastapi.responses import JSONResponse

from src.api.exceptions import ApiFailedConnectDataBase, ApiFailedToInsertBook
from src.api.model.books import BooksModels
from src.db.database import PostgresConnection
from src.db.repository.book import BookRepository

create = APIRouter()


@create.post("/", response_model=BooksModels)
async def create_handler(book: BooksModels) -> any:
    pg = None
    try:
        pg = PostgresConnection()
        pg._connect_db()
        repository = BookRepository(pg.conn)
        if not repository.create_book(book):
            raise ApiFailedToInsertBook
    except ApiFailedConnectDataBase:
        return HTTPException(status_code=500, detail="can't connect to database")
    except ApiFailedToInsertBook:
        return HTTPException(status_code=500, detail="can't insert to database")
    finally:
        if pg is not None:
            pg._close_connection()
    return JSONResponse(content=book, status_code=201)
