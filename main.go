package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	port := flag.String("port", "8080", "port to serve on")
	directory := flag.String("dir", "/data", "the static directory to host")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*directory)))
	fmt.Printf("serving %s on HTTP port: %s\n", *directory, *port)
	address := fmt.Sprintf(":%s", *port)

	err := http.ListenAndServe(address, nil)
	if err != nil {
		panic(err)
	}
}
