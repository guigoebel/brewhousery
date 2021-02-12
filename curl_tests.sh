# Read all
curl localhost:8080 | jq

# Post
curl localhost:8080 -d '{"name": "Tea", "description": "A nice cup of tea!", "price": 2.0, "sku": "TEA-WAT-MIL-SUG"}' | jq
curl localhost:8080 | jq

# Update
curl localhost:8080/3 -XPUT -d '{"name": "Mocha", "description": "Chocolate-flavoured variant of latte", "price": 3.00, "sku": "COF-MOC-VAR-LAT"}' | jq
curl localhost:8080 | jq

# Read one
curl localhost:8080/1 | jq
curl localhost:8080/2 | jq
curl localhost:8080/3 | jq