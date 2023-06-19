from pydantic import BaseModel, validator


class BooksModels(BaseModel):
    name: str
    price: float
    author: str

    @validator('name')
    def name_validation(cls, value):
        if not value:
            raise ValueError("name is empty")
        elif len(value) <= 3:
            raise ValueError(
                "name is too short, it has to be greater than 3 chars")
        return value.strip()

    @validator('price')
    def price_validation(cls, value):
        if value <= 0:
            raise ValueError("price value has to be greater than 0")
        if not isinstance(value, float):
            raise ValueError("price must to be float instance")
        return value

    @validator('author')
    def author_validation(cls, value):
        if not value:
            raise ValueError("author is empty")
        elif len(value) <= 3:
            raise ValueError(
                "author is too short, it has to be greater than 3 chars")
        return value.strip()
