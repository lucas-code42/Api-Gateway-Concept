from fastapi import APIRouter
from src.api.model.books import BooksModels

update = APIRouter()


@update.put("/")
async def update_handler():
    return
