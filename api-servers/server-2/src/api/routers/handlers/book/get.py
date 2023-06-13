from fastapi import APIRouter

get_book = APIRouter()


@get_book.get("/book")
async def get_book_by_id():
    return
