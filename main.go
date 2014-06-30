package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/go-martini/martini"
)

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!\n"
	})
	go m.Run()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello world!\n")
	})

	// Use the default middleware.
	n := negroni.Classic()
	// ... Add any other middlware here

	// add the router as the last handler in the stack
	n.UseHandler(mux)
	n.Run(":3001")
}