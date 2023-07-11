import datetime
import jwt
from jwt import InvalidAudienceError, ExpiredSignatureError
from fastapi import Header

from src.api.settings import settings


def generate_jwt_token() -> str:
    payload = {
        "aud": ["usr:lcs42", "pythonjwt"],
        "exp": datetime.datetime.now(tz=datetime.timezone.utc) + datetime.timedelta(seconds=900)
    }
    return jwt.encode(
        payload=payload,
        key=settings.JWT_KEY,
        algorithm=settings.JWT_DEFAULT_ALGORITHM,
    )


def decode_jwt_token_iss(token: str = Header()) -> bool:
    result = None
    print(token)
    try:
        # we can pass a list and if one of those are true the decode is successfully
        if decode := jwt.decode(
            token,
            key=settings.JWT_KEY,
            audience=["usr:lcs42", "pythonjwt"],  # use envs...
            algorithms=settings.JWT_DEFAULT_ALGORITHM
        ):
            result = True
    except InvalidAudienceError:
        raise InvalidAudienceError("Invalid audience")
    except ExpiredSignatureError:
        raise ExpiredSignatureError("Expired time")
    return result

def auth():
    def decorator(func):
        async def wrapper(*args, **kwargs):
            authorization = kwargs["authorization"]
            token = authorization.headers.get("authorization")
            if decode_jwt_token_iss(token):
                result = await func(*args, **kwargs)
                return result
        return wrapper
    return decorator