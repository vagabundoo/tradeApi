from fastapi import FastAPI

app = FastAPI()


@app.get("/trades")
async def root():
    return {"message": "Hello World"}

@app.get("/trades/filter")
async def root():
    return {"message": "Hello World"}

@app.post("/trades")
async def root():
    return {"message": "Hello World"}

@app.patch("/trades")
async def root():
    return {"message": "Hello World"}
