from fastapi import APIRouter

update_book = APIRouter()


@update_book.put("/book")
async def update():
    return
