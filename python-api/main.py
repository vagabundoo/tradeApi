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

trades[0].country = "USA"
print(trades[0])

@app.get("/trades")
async def get_trades(skip: int = 0, limit: int = 100):
    return trades[skip : skip + limit]

@app.get("/trades/filter")
async def filter_trades(
    id: int = None,
    country: str = None,
    year: int = None,
    commodity_code: str = None,
    flow: str = None
):
    filtered_trades = []

    for i in range(len(trades)):
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
     
    return filtered_trades

@app.post("/trades")
async def create_trade(trade: Trade):
    return trade

@app.patch("/trades/{id}")
async def patch_trade_by_id(
    id: int,
    country: str = None,
    year: int = None,
    commodity_code: str = None,
    flow: str = None
):
    for i in range(len(trades)):
        if id == trades[i].id: 
            if country: trades[i].country = country
            if year: trades[i].year = year
            if commodity_code: trades[i].commodity_code = commodity_code
            if flow: trades[i].flow = flow
            return trades[i]

    
