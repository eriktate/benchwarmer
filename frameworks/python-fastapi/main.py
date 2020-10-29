import os

from fastapi import FastAPI
from fastapi.responses import PlainTextResponse
from pydantic import BaseModel
import uvicorn

class JSONReq(BaseModel):
    greeting: str
    name: str

app = FastAPI()

@app.get("/hello", response_class=PlainTextResponse)
async def hello_handler():
    return "Hello, World!"

@app.post("/json")
async def json_handler(req: JSONReq):
    return {"msg": f"{req.greeting} {req.name}"}

if __name__ == "__main__":
    port = int(os.getenv("BENCH_PORT"))
    workers = int(os.getenv("BENCH_WORKERS"))
    print(f"Ports: {port} Workers: {workers}")

    uvicorn.run("main:app", port=port, host=os.getenv("BENCH_HOST"), access_log=False, workers=workers)
