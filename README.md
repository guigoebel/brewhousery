[![Build Status](https://travis-ci.com/saurabmish/Coffee-Shop.svg?branch=master)](https://travis-ci.com/saurabmish/Coffee-Shop)

# Coffee Shop

REST API for an online coffee shop

+ Commonly used HTTP methods
+ JSON request and response using curl

## Data Type

**Temporary in-memory storage**

```
{
    "ID": unique identifier                        integer
    "Name": type of coffee (latte, expresso, etc.) string
    "Description": brief info                      string
    "Price": USD                                   float
    "SKU": barcode,                                string
    "Manufactured On": UTC date                    string
    "Expires On": UTC date                         string
}
```

## Operations

+ Run the Go server from the project root.

  `go run server.go`

  **NOTE:** For better formatting, pipe curl operations with `jq`

+ HTTP verbs

  + GET request

    `curl localhost:8080 | jq`

  + POST (ID is auto-incremented)

    `curl localhost:8080 -d '{"name": "Tea", "description": "A nice cup of tea!"}' | jq`

  + PUT (replace data with ID 3)
   
    `curl localhost:8080/3 -XPUT -d '{"name": "Mocha", "description": "Chocolate-flavoured variant of latte"}' | jq`
