from fastapi import APIRouter

delete = APIRouter()


@delete.delete("/")
async def delete_book_by_id_handler():
    return


@delete.delete("/all")
async def delete_all():
    return
