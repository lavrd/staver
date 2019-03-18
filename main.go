package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var (
		port = flag.Int("p", 7777, "Set staver port")
		dist = flag.String("d", "./dist/", "Set staver dist")
	)

	flag.Parse()

	fs := http.FileServer(http.Dir(*dist))

	log.Println("Staver")
	log.Println("Dist folder:", *dist)
	log.Println("Port:", *port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", *port), logger(fs)))
}
