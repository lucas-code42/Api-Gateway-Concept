from pydantic import BaseModel, ValidationError, validator


class StockModels(BaseModel):
    books_id: int
    quantity: int

    @validator('books_id')
    def books_id_validation(cls, value):
        if not value:
            raise ValidationError("books_id is empty")
        elif value <= 0:
            raise ValidationError("books_id it has to be greater than 0")
        elif not isinstance(value, int):
            raise ValidationError("books_id it has to be a int instance")
        return value

    @validator('quantity')
    def quantity_validation(cls, value):
        if not value:
            raise ValidationError("quantity is empty")
        elif value <= 0:
            raise ValidationError("quantity it has to be greater than 0")
        elif not isinstance(value, int):
            raise ValidationError("quantity it has to be a int instance")
        return value
