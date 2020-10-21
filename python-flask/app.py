import json
import os

from flask import (Flask, request)

app = Flask(__name__)

@app.route("/hello", methods=["GET"])
def hello_handler():
    return "Hello, world!"

@app.route("/json", methods=["POST"])
def json_handler():
    req = json.loads(request.data)
    return f'{req["greeting"]} {req["name"]}'

if __name__ == "__main__":
    app.run(host=os.getenv("BENCH_HOST"), port=os.getenv("BENCH_PORT"))
