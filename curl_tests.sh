# Read all
curl localhost:8080/coffee/get/all | jq

# Post
curl localhost:8080/coffee/add -d '{"name": "Tea", "description": "A nice cup of tea!", "price": 2.0, "sku": "TEA-WAT-MIL-SUG"}' | jq
curl localhost:8080/coffee/get/all | jq

# Update
curl localhost:8080/coffee/modify/3 -XPUT -d '{"name": "Mocha", "description": "Chocolate-flavoured variant of latte", "price": 3.00, "sku": "COF-MOC-VAR-LAT"}' | jq
curl localhost:8080/coffee/get/all | jq

# Read one
curl localhost:8080/coffee/get/1 | jq
curl localhost:8080/coffee/get/2 | jq
curl localhost:8080/coffee/get/3 | jq

# Delete
curl localhost:8080/coffee/remove/2 -XDELETE
curl localhost:8080/coffee/get/all | jq
