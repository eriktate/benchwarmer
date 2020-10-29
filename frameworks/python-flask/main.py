import json
import os
import logging

from flask import (Flask, request)

app = Flask(__name__)

# disable logging so it doesn't interfere with benchmarks
log = logging.getLogger('werkzeug')
log.disabled = True

@app.route("/hello", methods=["GET"])
def hello_handler():
    return "Hello, World!"

@app.route("/json", methods=["POST"])
def json_handler():
    req = json.loads(request.data)
    return json.dumps({"msg": f'{req["greeting"]} {req["name"]}'})

if __name__ == "__main__":
    app.run(host=os.getenv("BENCH_HOST"), port=os.getenv("BENCH_PORT"))
