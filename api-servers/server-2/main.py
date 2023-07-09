from fastapi import FastAPI
from typing import Dict
from datetime import datetime
from src.api.settings import settings
from src.api.routers.routers import endpoints


app = FastAPI(
    description=settings.API_DESCRIPTION,
    version=settings.API_VERSION
)
app.include_router(endpoints)


@app.get("/")
async def health() -> Dict:
    return {"health": f"{datetime.now()}"}

if __name__ == "__main__":
    import uvicorn

    uvicorn.run(
        app="main:app",
        port=settings.API_PORT,
        log_level="info",
        use_colors=True
    )
