import jwt
from jwt import InvalidAudienceError

from src.api.settings import settings


def generate_jwt_token() -> str:
    payload = {"aud": ["urn:lcs42", "pythonjwt"]}
    return jwt.encode(
        payload=payload,
        key=settings.JWT_KEY,
        algorithm=settings.JWT_DEFAULT_ALGORITHM
    )


def decode_jwt_token_iss(token: str) -> bool:
    result = None
    try:
        # we can pass a list and if one of those are true the decode is successfully
        if decode := jwt.decode(
            token,
            key=settings.API_JWT_KEY,
            audience=["urn:qw", "332"],  # use envs...
            algorithms=settings.JWT_DEFAULT_ALGORITHM
        ):
            result = True
    except InvalidAudienceError:
        raise InvalidAudienceError("Invalid audience")
    return result
