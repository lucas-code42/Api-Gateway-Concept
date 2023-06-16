from fastapi import APIRouter

from src.api.routers.handlers.book.create import create
from src.api.routers.handlers.book.delete import delete
from src.api.routers.handlers.book.get import get
from src.api.routers.handlers.book.update import update


endpoints = APIRouter()
endpoints.include_router(router=create, prefix="/book")
endpoints.include_router(router=delete, prefix="/book")
endpoints.include_router(router=get, prefix="/book")
endpoints.include_router(router=update, prefix="/book")
