package main

import (
	"fmt"
	"net/http"

	"github.com/Zemavong/Go-React-Todo/router"
)

func main() {
	r := router.Router()
	fmt.Println("starting the server on port 9000...")

	http.ListenAndServe(":9000", r)
}
