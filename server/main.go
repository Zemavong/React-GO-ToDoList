package main

import (
	"fmt"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Printl("starting the server on port 9000...")

	http.ListenAndServe(":9000", r)
}
