package main

import (
	"fmt"
	"github.com/72sevenzy2/golang-API/internal/routes"
	"github.com/72sevenzy2/golang-API/internal/service"
	"net/http"
)

func main() {
	service := &service.GreetCounter{}

	r := routes.NewRouter(service)

	fmt.Println("server running on port 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}