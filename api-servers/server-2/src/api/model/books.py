from pydantic import BaseModel, ValidationError, validator


class BooksModels(BaseModel):
    name: str
    price: float
    author: str

    @validator('name')
    def name_validation(cls, value):
        if not value:
            raise ValidationError("name is empty")
        elif len(value) <= 3:
            raise ValidationError(
                "name is too short, it has to be greater than 3 chars")
        return value.strip()

    @validator('price')
    def price_validation(cls, value):
        if value <= 0:
            raise ValidationError("price value has to be greater than 0")
        if not isinstance(value, float):
            raise ValidationError("price must to be float instance")
        return value

    @validator('author')
    def author_validation(cls, value):
        if not value:
            raise ValidationError("author is empty")
        elif len(value) <= 3:
            raise ValidationError(
                "author is too short, it has to be greater than 3 chars")
        return value.strip()
