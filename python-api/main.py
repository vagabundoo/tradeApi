from typing import Annotated, Literal

from fastapi import FastAPI, Query
from pydantic import BaseModel

app = FastAPI()

class Trade(BaseModel):
    id: int
    country: str
    year: int
    commodity_code: str
    commodity_desc: str 
    flow: str
    trade_usd: int

# dummy data
trades = [
    Trade(id=1, country="Afghanistan", year=2016, commodity_code="010410", commodity_desc="Sheep, live", flow="Export", trade_usd=6088),
	Trade(id=2, country="Argentina", year=1996, commodity_code="010410", commodity_desc="Sheep, live", flow="Import", trade_usd=798372),
	Trade(id=3, country="Australia", year=2007, commodity_code="010410", commodity_desc="Sheep, live", flow="Export", trade_usd=219992664)
]

@app.get("/trades")
async def get_trades(skip: int = 0, limit: int = 100):
    return trades[skip : skip + limit]

@app.get("/trades/filter")
async def filter_trades(
    id: int | None = None,
    country: str | None = None,
    year: int | None = None,
    commodity_code: str | None = None,
    flow: str | None = None
):
    filtered_trades = []

    for i in trades:
        # we iterate through records, and continue early if one of the filter criteria is not met
        if id: 
            if id != trades[i].id: continue
        if country: 
            if country != trades[i].country: continue
        if year: 
            if year != trades[i].year: continue
        if commodity_code: 
            if commodity_code != trades[i].commodity_code: continue
        if flow: 
            if flow != trades[i].flow: continue

        filtered_trades.append(trades[i])

        

    return {"message": "Hello World"}

@app.post("/trades")
async def root():
    return {"message": "Hello World"}

@app.patch("/trades/{id}")
async def root():
    return {"message": "Hello World"}
