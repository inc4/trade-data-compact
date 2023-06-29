package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/inc4/go-template/pkg/api"
)

func main() {
	// Basic command-line flag parsing
	port := flag.Int("port", 8080, "example message for port number")
	exampleStr := flag.String("example", "foo", "help message")
	flag.Parse()

	fmt.Println("command-line:", *port, *exampleStr)

	http.HandleFunc("/hello", api.HandleHello)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
