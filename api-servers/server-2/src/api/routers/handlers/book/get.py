from fastapi.encoders import jsonable_encoder
from fastapi.responses import JSONResponse
from fastapi import APIRouter, HTTPException
from api.exceptions import ApiFailedConnectDataBase, ApiFailedToGetBookById
from src.db.database import PostgresConnection
from src.db.repository.book import BookRepository


get = APIRouter()


@get.get("/")
async def get_book(book_id: int):
    pg = None
    result = None
    try:
        pg = PostgresConnection()
        pg._connect_db()
        repository = BookRepository(pg.conn)
        result = repository.get_book_by_id(book_id)
        if result is None:
            raise
    except ApiFailedConnectDataBase:
        raise HTTPException(
            status_code=500, detail="can't connect to database")
    except ApiFailedToGetBookById:
        raise HTTPException(status_code=500, detail="can't get book by id")
    finally:
        if pg.conn is not None:
            pg._close_connection()
    return JSONResponse(content=jsonable_encoder(result), status_code=200)


@get.get("/all")
async def get_all():
    return
