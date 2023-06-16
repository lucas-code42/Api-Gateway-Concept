from fastapi import APIRouter

create = APIRouter()


@create.post("/")
async def create_handler():
    return
