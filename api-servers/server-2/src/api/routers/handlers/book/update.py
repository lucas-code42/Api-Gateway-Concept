from fastapi import APIRouter

update = APIRouter()


@update.put("/")
async def update_handler():
    return
