import datetime
import time
import jwt

# jwt_payload = jwt.encode(
#     {"exp": datetime.datetime.now(tz=datetime.timezone.utc) + datetime.timedelta(seconds=30)},
#     "secret",
# )

# time.sleep(3)

# # JWT payload is now expired
# # But with some leeway, it will still validate
# print(jwt.decode(jwt_payload, "secret", leeway=10, algorithms=["HS256"]))


def generate_jwt_token() -> str:
    payload = {
        "aud": ["urn:lcs42", "pythonjwt"],
        "exp": datetime.datetime.now(tz=datetime.timezone.utc) + datetime.timedelta(seconds=5)
    }
    return jwt.encode(
        payload=payload,
        key="teste",
        algorithm="HS256"
    )


def decode_jwt_token_iss(token: str) -> bool:
    result = None
    try:
        # we can pass a list and if one of those are true the decode is successfully
        if decode := jwt.decode(
            token,
            key="teste",
            audience=["urn:lcs41", "pythonjwt"],  # use envs...
            algorithms=["HS256"]
        ):
            result = True
    except jwt.InvalidAudienceError:
        raise jwt.InvalidAudienceError("Signature error")
    except jwt.ExpiredSignatureError:
        raise jwt.InvalidAudienceError("Expired time")
    return result


# jwt_code = generate_jwt_token()
# time.sleep(1)
# print(decode_jwt_token_iss(jwt_code))
# print(jwt_code)


def auth():
    def decorator(func):
        def wrapper(**kwargs):
            token = kwargs["token"]
            if decode_jwt_token_iss(token):
                result = func(**kwargs)
                return result
        return wrapper
    return decorator


@auth()
def api(token: str, data: dict) -> dict:
    for i in range(10):
        print(i)
        print(token.__class__)

    data.update({"teste": "oi"})
    return data


jwt_code = generate_jwt_token()
api(token=jwt_code, data={"data":10})


# def meu_decorator(comentario):
#     def decorator(funcao):
#         def wrapper(*args, **kwargs):
#             print(comentario)
#             resultado = funcao(*args, **kwargs)
#             print("funcao exc aqui ", resultado)
#             return resultado
#         return wrapper
#     return decorator


# @meu_decorator(comentario="Essa é uma função decorada.")
# def minha_funcao(x, y):
#     return x + y


# minha_funcao(1, 2)
