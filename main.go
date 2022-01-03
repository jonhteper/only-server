package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("Use only-server <port> <local-dir>")
	}
	port := args[0]

	fmt.Printf("Listen server in http://localhost:%v ...", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), http.FileServer(http.Dir(args[1]))))
}
