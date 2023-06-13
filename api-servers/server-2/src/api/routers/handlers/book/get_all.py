from fastapi import APIRouter

get_all_books = APIRouter()


@get_all_books.get("/books")
async def get_all():
    return
