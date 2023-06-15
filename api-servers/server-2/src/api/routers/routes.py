from fastapi import APIRouter

from src.api.routers.handlers.book.create import create_book
from src.api.routers.handlers.book.delete import delete_book
from src.api.routers.handlers.book.get_all import get_all_books
from src.api.routers.handlers.book.get import get_book
from src.api.routers.handlers.book.update import update_book


endpoints = APIRouter()
endpoints.include_router(router=create_book, prefix="/book")
endpoints.include_router(router=delete_book, prefix="/book")
endpoints.include_router(router=get_all_books, prefix="/books")
endpoints.include_router(router=get_book, prefix="/book")
endpoints.include_router(router=update_book, prefix="/book")
