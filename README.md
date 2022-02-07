[![Build Status](https://app.travis-ci.com/saurabh-mish/brewhousery.svg?branch=master)](https://app.travis-ci.com/saurabh-mish/brewhousery)
[![codecov](https://codecov.io/gh/saurabmish/brewhousery/branch/master/graph/badge.svg?token=YUPUN62OPY)](https://codecov.io/gh/saurabmish/brewhousery)

# Brewhousery

APIs of commonly used HTTP verbs GET, POST, PUT, and DELETE are implemented in a REST-ful manner using micro-services in Go.

## Features

+ Retrieve single product

+ Retrieve all products

+ Add a new product

+ Modify an existing product

+ Delete an existing product

## Data Type

```
{
    ID:              unique integer identifier for coffee product
    Name:            name of coffee (latte, espresso, etc.)
    Description:     brief information about coffee
    Price:           floating point values in USD
    SKU:             stock-keeping unit is a string separated by "-" after every third character
    Manufactured On: string in UTC format
    Expires On:      string in UTC format
}
```

## Operations

**NOTE**: Execute commands from project root

+ Run test cases (in verbose mode):

  `go test -v`

+ Run Go server:

  `go run server.go`

+ Open a new terminal tab and:

  1. Get specific product:

     `curl localhost:8080/coffee/get/2 | jq`

  2. Get all products:

     `curl localhost:8080/coffee/get/all | jq`

  3. Create a new product:

     `curl localhost:8080/coffee/add -d '{"name": "Tea", "description": "A nice cup of tea!", "price": 2.0, "sku": "TEA-WAT-MIL-SUG"}' | jq`

  4. Replace an existing product:

     `curl localhost:8080/coffee/modify/3 -XPUT -d '{"name": "Mocha", "description": "Chocolate-flavoured variant of latte", "price": 3.00, "sku": "COF-MOC-VAR-LAT"}' | jq`

  5. Delete an existing product

     `curl localhost:8080/coffee/remove/2 -XDELETE`
