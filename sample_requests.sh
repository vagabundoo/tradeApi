#curl http://localhost:8080/trades
echo filter by year
curl "http://localhost:8080/trades/filter?year=2016"

echo filter by multiple fields
curl "http://localhost:8080/trades/filter?country=Australia&commodity_code=010410"

curl http://localhost:8080/trades \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"ID": 4, "Country": "United States","Year": 2007,"Commodity_code": "010410","Commodity_desc": "Sheep, live","Flow": "Export","Trade_usd": 219992664}'




    