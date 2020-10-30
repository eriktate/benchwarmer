import os
from uuid import uuid4

from typing import Optional
from fastapi import FastAPI
from fastapi.responses import PlainTextResponse
from pydantic import BaseModel
from datetime import datetime
from psycopg2 import pool
import uvicorn

host = os.getenv("BENCH_DB_HOST")
port = os.getenv("BENCH_DB_PORT")
user = os.getenv("BENCH_DB_USER")
passwd = os.getenv("BENCH_DB_PASSWORD")

db = pool.SimpleConnectionPool(
    1,
    50,
    host=host,
    port=port,
    user=user,
    password=passwd,
)

createQuery = """
INSERT INTO greetings (id, greeting, name)
VALUES (%s, %s, %s);
"""

getQuery = """
SELECT
	id,
	greeting,
	name,
	used,
	created_at,
	updated_at
FROM
	greetings
WHERE
	id = %s;
"""
class Greeting(BaseModel):
    id: Optional[str]
    greeting: str
    name: str
    used: Optional[bool]
    createdAt: Optional[datetime]
    updatedAt: Optional[datetime]

app = FastAPI()

@app.get("/hello", response_class=PlainTextResponse)
async def hello_handler():
    return "Hello, World!"

@app.post("/json")
async def json_handler(req: Greeting):
    return {"msg": f"{req.greeting} {req.name}"}

@app.post("/db")
async def db_handler(req: Greeting):
    print("Handling DB!")
    req.id = str(uuid4())
    conn = db.getconn()
    cur = conn.cursor()
    cur.execute(createQuery, (req.id, req.greeting, req.name))
    print("Created greeting!")
    cur.execute(getQuery, (req.id,))
    print("Fetched greeting!")
    row = cur.fetchone()

    # generate response
    res = Greeting(greeting=row[1], name=row[2])
    res.id = row[0]
    res.used = row[3]
    res.createdAt = row[4]
    res.updatedAt = row[5]

    # clean up
    conn.commit()
    cur.close()
    conn.close()

    return res

if __name__ == "__main__":
    port = int(os.getenv("BENCH_PORT"))
    workers = int(os.getenv("BENCH_WORKERS"))
    print(f"Ports: {port} Workers: {workers}")

    uvicorn.run("main:app", port=port, host=os.getenv("BENCH_HOST"), access_log=False, workers=workers)
