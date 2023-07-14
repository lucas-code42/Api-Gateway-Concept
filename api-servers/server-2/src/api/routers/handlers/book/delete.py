from fastapi import APIRouter, HTTPException, Response, Depends

from src.db.database import PostgresConnection
from src.db.repository.book import BookRepository
from src.api.exceptions import ApiFailedToInsertBook, ApiFailedConnectDataBase

from src.api.security.jwt_token.auth import decode_jwt_token_iss

delete = APIRouter()


@delete.delete("/", dependencies=[Depends(decode_jwt_token_iss)])
async def delete_book_by_id_handler(book_id: int):
    pg = None
    try:
        pg = PostgresConnection()
        await pg._connect_db()
        repository = BookRepository(pg.conn)
        if not await repository.delete_book_by_id(book_id):
            raise ApiFailedToInsertBook
    except ApiFailedConnectDataBase:
        raise HTTPException(
            status_code=500, detail="can't connect to database")
    except ApiFailedToInsertBook:
        raise HTTPException(
            status_code=500, detail="can't remove from database")
    finally:
        if pg.conn is not None:
            await pg._close_connection()
    return Response(status_code=204)
