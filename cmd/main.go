package main

import (
	db "github.com/adam-fraga/adhora/db"
	h "github.com/adam-fraga/adhora/handlers"
	r "github.com/adam-fraga/adhora/routers"
)

func main() {

	db := db.DB{}
	db.NewConnection()
	defer db.Close()
	db.Ping()

	router := r.NewRouter()
	router.ServeStatic()

	//Should be in Routes
	router.HandleFunc("/", h.IndexHandler)
	router.HandleFunc("/contact", h.ContactHandler)
	router.HandleFunc("/about", h.AboutHandler)
	router.HandleFunc("/services", h.ServicesHandler)
	router.HandleFunc("/login", h.SigninHandler)
	router.HandleFunc("/register", h.SignupHandler)
	router.HandleFunc("/dashboard", h.DashboardHandler)

	router.ListenAndServe(":3000")
}
