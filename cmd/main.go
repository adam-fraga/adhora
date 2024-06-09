package main

import (
	h "github.com/adam-fraga/adhora/handlers"
	r "github.com/adam-fraga/adhora/routers"
)

func main() {

	router := r.NewRouter()
	router.ServeStatic()

	//Should be in Routes
	router.HandleFunc("/", h.IndexHandler)
	router.HandleFunc("/contact", h.ContactHandler)
	router.HandleFunc("/about", h.AboutHandler)

	router.ListenAndServe(":3000")
}
