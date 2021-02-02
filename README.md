# Car-Inventory

REST API for car inventory operations

+ Native packages, no dependencies
+ Commonly used HTTP methods
+ JSON request and response using curl

## Features

+ Retrieve all cars

+ Retrieve details of a specific car

+ Add a new car

+ Get a car for spot allocation

## Methods

+ `GET /cars`

  Returns all cars

+ `GET /cars/{id}`

  Returns details of a specific car

+ `POST /cars`

  Add a new car to inventory

+ `GET /cars/random`

  Get a random car from inventory

## Data Type

**Temporary in-memory storage**

```
{
    "id": "UUID",
    "name": "Brand name",
    "kind": "Gas / Hybrid / Electric",
    "manufacturer": "Name of manufacturer",
    "license": "License number",
    "available": "Yes or No"
}
```
