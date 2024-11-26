package main

import (
	"Checkmarx/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.NewRouter()
	fmt.Println("Server is running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
