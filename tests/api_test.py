import requests
import os

base_url = f'http://{os.getenv("HOST_ADDR")}'
print(f"Testing {base_url}")


def test_hello():
    res = requests.get(f"{base_url}/hello")
    assert(res.text == "Hello, World!")

def test_json():
    res = requests.post(f"{base_url}/json", json={"greeting": "Hello", "name": "Benchwarmer"})
    data = res.json()
    assert(data["msg"] == "Hello Benchwarmer")
