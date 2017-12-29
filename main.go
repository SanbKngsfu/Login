package main

import (
	"log"
	"login/route"
	"net/http"
)

func main() {
	route.Route()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("ListenAndServe", err.Error())
	}
}
