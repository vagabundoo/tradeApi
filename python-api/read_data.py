import polars as pl
import os
print(os.getcwd())

trade_schema = pl.Schema({
    "country_or_area": pl.String, 
    "year": pl.Int16, 
    "comm_code": pl.String, 
    "commodity": pl.String, 
    "flow": pl.String, 
    "trade_usd": pl.Int64
})


trade_df = pl.read_csv(
    "../data/commodity_trade_statistics_data.csv",
    schema = trade_schema,
    n_rows = 100000,
    truncate_ragged_lines=True
)
print(trade_df)
"""
trade_df = pl.scan_csv(
    "../data/commodity_trade_statistics_data.csv"
).select(v
).filter(
    pl.col("year") > 2015
)
"""



