package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")
	fmt.Printf("Hello world, got port *:%s", port)

	http.HandleFunc("/cron", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world")
	})
	fmt.Printf("running cron job at port *:%s", port)
	http.ListenAndServe(":"+port, nil)

}
