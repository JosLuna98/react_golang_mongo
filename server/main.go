package main

import (
	"fmt"
	"log"
	"net/http"

	"server/router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}

//! Uncoment it to testing
// import (
// 	"server/test"
// )

// func main() {
// 	main_testing.Test()
// }
