# Go-REST-API-with-MongoDB

Simple REST API in GO Using the official MongoDB driver


### To get basic external modules for REST API

 ```sh
go get github.com/julienschmidt/httprouter
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/mongo/options
go get go.mongodb.org/mongo-driver/bson
```

* [httprouter](github.com/julienschmidt/httprouter) - HTTP request router
* [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) - official MongoDB Go driver


## To run the server
 ```sh
* go run .
```

## Routing
 Routing | HTTP | Description
 --------|------|------------
/users | GET | GetUsers
/users/:id | GET | GetUser
/users | POST | AddUser
/users/:id | PUT | UpdateUser
/users/:id | DELETE | DeleteUser
