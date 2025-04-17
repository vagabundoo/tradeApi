curl http://localhost:8080/trades \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"ID": 4, "Country": "United States","Year": 2007,"Commodity_code": "010410","Commodity_desc": "Sheep, live","Flow": "Export","Trade_usd": 219992664}'

#curl http://localhost:8080/trades

curl http://localhost:8080/trades/2007

    