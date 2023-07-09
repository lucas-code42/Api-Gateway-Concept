from fastapi import APIRouter, HTTPException
from fastapi.responses import JSONResponse
from src.api.model.books import BooksModels
from src.db.database import PostgresConnection
from src.db.repository.book import BookRepository
from src.api.exceptions import ApiFailedConnectDataBase, ApiFailedToVerifyId, ApiFailedToUpdateBooks
from fastapi.encoders import jsonable_encoder

update = APIRouter()


@update.put("/", response_model=BooksModels)
async def update_handler(book: BooksModels):
    pg = None
    result = None
    try:
        pg = PostgresConnection()
        await pg._connect_db()
        repository = BookRepository(pg.conn)
        result = await repository.udpate_book_by_id(book)
    except ApiFailedConnectDataBase:
        raise HTTPException(
            status_code=500, detail="can't connect to database")
    except ApiFailedToVerifyId:
        raise HTTPException(
            status_code=500, detail="book id not exist")
    except ApiFailedToUpdateBooks:
        raise HTTPException(
            status_code=500, detail="can't update book")
    finally:
        if pg.conn is not None:
            await pg._close_connection()
    return JSONResponse(content=jsonable_encoder(result), status_code=201)
