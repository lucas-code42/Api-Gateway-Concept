from fastapi import APIRouter, HTTPException, Response

from src.db.database import PostgresConnection
from src.db.repository.book import BookRepository
from src.api.exceptions import ApiFailedToInsertBook, ApiFailedConnectDataBase

delete = APIRouter()


@delete.delete("/")
async def delete_book_by_id_handler(book_id: int):
    pg = None
    try:
        pg = PostgresConnection()
        pg._connect_db()
        repository = BookRepository(pg.conn)
        if not repository.delete_book_by_id(book_id):
            raise ApiFailedToInsertBook
    except ApiFailedConnectDataBase:
        raise HTTPException(
            status_code=500, detail="can't connect to database")
    except ApiFailedToInsertBook:
        raise HTTPException(
            status_code=500, detail="can't remove from database")
    finally:
        if pg.conn is not None:
            pg._close_connection()
    return Response(status_code=204)


@delete.delete("/all")
async def delete_all():
    return
