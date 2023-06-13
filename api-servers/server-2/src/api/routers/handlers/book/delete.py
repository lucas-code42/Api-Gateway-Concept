from fastapi import APIRouter

delete_book = APIRouter()


@delete_book.delete("/book")
async def delete_book_by_id():
    return
