package main

import (
	"fmt"
	"net/http"
	"os"
	
	"github.com/julienschmidt/httprouter"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	fmt.Printf("Hello world, got port *:%s", port)
	
	// We are using a third-party library here to demonstrate an example of using package in aws beanstalk
	r := httprouter.New()
	r.POST("/cron", func(w http.ResponseWriter, r *http.Request) {
		// Carry out the cron job
		log.Println("Running cron")
		fmt.Fprint(w, "Hello world")
	})
	
	log.Printf("Listening to server at port *:%s\n", port)
	http.ListenAndServe(":"+port, r)
}
