from fastapi import APIRouter
from pydantic import BaseModel
from src.api.security.jwt_token.token import generate_jwt_token

jwt = APIRouter()


class JwtToken(BaseModel):
    token: str


@jwt.get("/", response_model=JwtToken)
def delivery_token():
    return JwtToken(**{"token": generate_jwt_token()})
