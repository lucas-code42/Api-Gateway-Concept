import datetime
import jwt
from jwt import InvalidAudienceError, ExpiredSignatureError

from src.api.settings import settings


def generate_jwt_token() -> str:
    payload = {
        "aud": ["usr:lcs42", "pythonjwt"],
        "exp": datetime.datetime.now(tz=datetime.timezone.utc) + datetime.timedelta(seconds=30)
    }
    return jwt.encode(
        payload=payload,
        key=settings.JWT_KEY,
        algorithm=settings.JWT_DEFAULT_ALGORITHM,
    )


def decode_jwt_token_iss(token: str) -> bool:
    result = None
    try:
        # we can pass a list and if one of those are true the decode is successfully
        if decode := jwt.decode(
            token,
            key=settings.API_JWT_KEY,
            audience=["usr:lcs42", "pythonjwt"],  # use envs...
            algorithms=settings.JWT_DEFAULT_ALGORITHM
        ):
            result = True
    except InvalidAudienceError:
        raise InvalidAudienceError("Invalid audience")
    except ExpiredSignatureError:
        raise ExpiredSignatureError("Expired time")
    return result
