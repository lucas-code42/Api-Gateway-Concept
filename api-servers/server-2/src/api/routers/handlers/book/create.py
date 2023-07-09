from fastapi import APIRouter, HTTPException
from fastapi.responses import JSONResponse
from fastapi.encoders import jsonable_encoder
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
        await pg._connect_db()
        repository = BookRepository(pg.conn)
        book_id = await repository.create_book(book)
        if not book_id:
            raise ApiFailedToInsertBook
    except ApiFailedConnectDataBase:
        raise HTTPException(
            status_code=500, detail="can't connect to database")
    except ApiFailedToInsertBook:
        raise HTTPException(status_code=500, detail="can't insert to database")
    finally:
        if pg.conn is not None:
            await pg._close_connection()
    return JSONResponse(content=jsonable_encoder(book), status_code=201)
