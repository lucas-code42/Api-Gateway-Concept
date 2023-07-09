from fastapi.encoders import jsonable_encoder
from fastapi.responses import JSONResponse
from fastapi import APIRouter, HTTPException
from src.api.exceptions import (
    ApiFailedConnectDataBase,
    ApiFailedToGetBookById,
    ApiFaliledToGetAllBooks
)
from src.db.database import PostgresConnection
from src.db.repository.book import BookRepository
from typing import List
from src.api.model.books import BooksModels


get = APIRouter()


@get.get("/")
async def get_book(book_id: int):
    pg = None
    result = None
    try:
        pg = PostgresConnection()
        await pg._connect_db()
        repository = BookRepository(pg.conn)
        result = await repository.get_book_by_id(book_id)
        if result is None:
            raise
    except ApiFailedConnectDataBase:
        raise HTTPException(
            status_code=500, detail="can't connect to database")
    except ApiFailedToGetBookById:
        raise HTTPException(status_code=500, detail="can't get book by id")
    finally:
        if pg.conn is not None:
            await pg._close_connection()
    return JSONResponse(content=jsonable_encoder(result), status_code=200)


@get.get("/all", response_model=List[BooksModels])
async def get_all():
    pg = None
    result = []
    try:
        pg = PostgresConnection()
        await pg._connect_db()
        reposittory = BookRepository(pg.conn)
        result = await reposittory.get_all_books()
        if not result:
            raise ApiFaliledToGetAllBooks
    except ApiFailedConnectDataBase:
        raise HTTPException(
            status_code=500, detail="can't connect to database")
    except ApiFaliledToGetAllBooks:
        raise HTTPException(
            status_code=500, detail="can't get all books")
    finally:
        if pg.conn is not None:
            await pg._close_connection()
    return JSONResponse(content=jsonable_encoder(result), status_code=200)
