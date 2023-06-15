from fastapi import FastAPI
from typing import Dict
from datetime import datetime
from src.api.settings import Settings
from src.api.routers.routes import endpoints



app = FastAPI(
    description=Settings.API_DESCRIPTION
)
app.include_router(endpoints)


@app.get("/")
async def health() -> Dict:
    return {"health": f"{datetime.now()}"}

if __name__ == "__main__":
    import uvicorn

    uvicorn.run(
        app="main:app",
        port=Settings.PORT,
        log_level="info"
    )
