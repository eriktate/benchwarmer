import os

from fastapi import FastAPI
from pydantic import BaseModel
import uvicorn

class JSONReq(BaseModel):
    greeting: str
    name: str

app = FastAPI()

@app.get("/hello")
async def hello_handler():
    return "Hello, world!"

@app.post("/json")
async def json_handler(req: JSONReq):
    return {"msg": f"{req.greeting} {req.name}"}

if __name__ == "__main__":
    uvicorn.run(app, port=int(os.getenv("BENCH_PORT")), host=os.getenv("BENCH_HOST"))
