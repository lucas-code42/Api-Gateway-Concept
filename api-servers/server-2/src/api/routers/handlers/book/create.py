from fastapi import APIRouter, HTTPException
from fastapi.responses import JSONResponse

from src.api.exceptions import ApiFailedConnectDataBase, ApiFailedToInsertBook
from src.api.model.books import BooksModels
from src.db.database import PostgresConnection
from src.db.repository.book import BookRepository


create = APIRouter()


@create.post("/", response_model=BooksModels)
async def create_handler(book: BooksModels):
    pg = None
    try:
        pg = PostgresConnection()
        pg._connect_db()
        repository = BookRepository(pg.conn)
        if not repository.create_book(book):
            raise ApiFailedToInsertBook
    except ApiFailedConnectDataBase:
        raise HTTPException(
            status_code=500, detail="can't connect to database")
    except ApiFailedToInsertBook:
        raise HTTPException(status_code=500, detail="can't insert to database")
    finally:
        if pg.conn is not None:
            pg._close_connection()
    return JSONResponse(content=book, status_code=201)
