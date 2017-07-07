package main

import (
	"fmt"
	"os"
)

func main() {

	port := os.Getenv("PORT")
	// fmt.Printf("Hello world, got port *:%s", port)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Hello world")
	// })
	// http.ListenAndServe(":"+port, nil)

	fmt.Printf("running cron job at port *:%s", port)
}
