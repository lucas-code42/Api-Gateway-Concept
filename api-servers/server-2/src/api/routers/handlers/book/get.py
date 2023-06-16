from fastapi import APIRouter

get = APIRouter()


@get.get("/")
async def get_book():
    return


@get.get("/all")
async def get_all():
    return
