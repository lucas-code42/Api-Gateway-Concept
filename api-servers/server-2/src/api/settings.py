from dotenv import dotenv_values
from src.api.exceptions import ApiFailedLoadEnvs


class Settings:
    """api setup"""

    API_DESCRIPTION: str = "CRUD-API for management virtual books"
    API_VERSION: str = "0.0.1"

    def __load_envs(self):
        self.envs: dict = dotenv_values(".env")
        if self.envs is None:
            raise ApiFailedLoadEnvs

    def __init__(self):
        self.__load_envs()

        # TODO ajustar if/else... não ficou muito legal está verboso. KISS
        jwt_default_algorithm = self.envs.get("JWT_DEFAULT_ALGORITHM")
        if jwt_default_algorithm:
            self.JWT_DEFAULT_ALGORITHM: str = jwt_default_algorithm
        else:
            self.JWT_DEFAULT_ALGORITHM: str = ""

        jwt_key = self.envs.get("JWT_KEY")
        if jwt_key:
            self.JWT_KEY: str = jwt_key
        else:
            self.JWT_KEY: str = ""

        port = self.envs.get("API_PORT")
        if port:
            self.API_PORT: int = int(port)
        else:
            self.API_PORT: int = 8000

        db_port = self.envs.get("DB_PORT")
        if db_port:
            self.DB_PORT: str = db_port
        else:
            self.DB_PORT: str = ""

        db_databse = self.envs.get("DB_DATABASE")
        if db_databse:
            self.DB_DATABSE: str = db_databse
        else:
            self.DB_DATABSE: str = ""

        user = self.envs.get("DB_USER")
        if user:
            self.DB_USER: str = user
        else:
            self.DB_USER: str = ""

        db_password = self.envs.get("DB_PASSWORD")
        if db_password:
            self.DB_PASSWORD: str = db_password
        else:
            self.DB_PASSWORD: str = ""


try:
    settings = Settings()
except ApiFailedLoadEnvs:
    print("Could not load env variables")
    exit(1)
