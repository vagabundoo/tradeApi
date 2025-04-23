import polars as pl
import os
print(os.getcwd())
trade_df = pl.read_csv(
    "../data/commodity_trade_statistics_data.csv",
    n_rows = 10000
)



print(trade_df)