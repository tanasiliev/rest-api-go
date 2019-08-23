package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Router struct {
	host string
	port int
}

func (r *Router) Int() {
	router := httprouter.New()

	router.GET("/users", GetUsers)

	router.GET("/users/:id", GetUser)

	router.POST("/users", AddUser)

	router.PUT("/users/:id", UpdateUser)

	router.DELETE("/users/:id", DeleteUser)

	serverAddr := r.host + ":" + fmt.Sprint(r.port)

	fmt.Println("The rest server runnig on http://" + serverAddr)

	log.Fatal(http.ListenAndServe(serverAddr, router))
}
