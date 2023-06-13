from fastapi import APIRouter

create = APIRouter(prefix="/book", tags=["create"])
get_by_id = APIRouter("/book", tags=["get by id"])
get_all = APIRouter("/books", tags=["get all books"])
update = APIRouter("/book", tags=["update by id"])
delete = APIRouter("/")
