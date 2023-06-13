from fastapi import APIRouter

create_book = APIRouter()


@create_book.post("/")
async def create():
    return
