package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"ungraded-challenge-5/config"
	"ungraded-challenge-5/handler"
)

func main() {
	router, server := config.SetupServer()
	db := &handler.NewRegisterHandler{DB: config.GetDatabase()}

	router.GET("/register", LoggingMiddleware(db.Register))
	router.GET("/login", LoggingMiddleware(db.Login))

	fmt.Println("Server running on port :8080")
	log.Fatal(server.ListenAndServe())
}

func LoggingMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		log.Printf("HTTP request sent to %s %s", r.Method, r.URL.Path)
		next(w, r, p)
	}
}
