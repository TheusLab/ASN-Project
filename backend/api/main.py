from fastapi import FastAPI
from search import search

app = FastAPI()

@app.get("/api/search")
def read_root(q: str):
    return search(q)