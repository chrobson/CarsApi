# REST API example aplication
This is a example of a basic application providing a REST API using gin-gonic package

The entire application is contained within the main.go file.

# Endpoints
Endpoints are described below.

## Get list of all cars
### Request
 `Get /cars/`
 
    curl -i -H 'Accept: application/json' http://localhost:8080/cars/
 
 ### Response
 
    HTTP/1.1 200 OK
    Content-Length: 251
    Content-Type: application/json; charset=utf-8
    Date: Mon, 24 Oct 2022 18:57:25 GMT
    [
    {
        "id": "1",
        "color": "Blue",
        "price": 25355.25
    },
    {
        "id": "2",
        "color": "Red",
        "price": 55122.99
    },
    {
        "id": "3",
        "color": "Green",
        "price": 62332.33
    }
]

## Create a new Car

### Request

`POST /cars/`

    curl -d '{"id":"4", "color":"Pink", "price":"89999.44"}' -H "Content-Type: application/json" -X POST http://localhost:8080/cars
    
 ## Get specific Car by his Id
 
 ### Request
 
 `GET /cars/id`
 
    curl -i -H 'Accept: application/json' http://localhost:8080/cars/2
    
 ## TODO
 * Integrate app with database (sqllite/postgres/mongodb)
 * Add auto-incrementator based on ID
 * Add message broker (RabbitMQ)
 * Add docker and manual how to deploy app
