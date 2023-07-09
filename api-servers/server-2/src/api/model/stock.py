from pydantic import BaseModel, validator
import deprecation
from src.api.settings import settings


@deprecation.deprecated(
    deprecated_in=settings.API_VERSION,
    removed_in="2023-07-08",
    details="I gave up doing :)"
)
class StockModels(BaseModel):
    books_id: int
    quantity: int

    @validator('books_id')
    def books_id_validation(cls, value):
        if not value:
            raise ValueError("books_id is empty")
        elif value <= 0:
            raise ValueError("books_id it has to be greater than 0")
        elif not isinstance(value, int):
            raise ValueError("books_id it has to be a int instance")
        return value

    @validator('quantity')
    def quantity_validation(cls, value):
        if not value:
            raise ValueError("quantity is empty")
        elif value <= 0:
            raise ValueError("quantity it has to be greater than 0")
        elif not isinstance(value, int):
            raise ValueError("quantity it has to be a int instance")
        return value
