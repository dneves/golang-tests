package main

import "fmt"
import "net/http"
import "./templates"

// generate templates: ~/go/bin/qtc

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func helpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", templates.Hello("Foo"))
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/help", helpHandler)
	http.ListenAndServe(":8080", nil)
}
