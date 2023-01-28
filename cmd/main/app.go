package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"testAPI/internal/user"
	"time"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	name := params.ByName("name")
	fmt.Fprintf(w, "Welcome %s!\n", name)
	fmt.Println(r.URL)
	fmt.Println(r.Header)
	fmt.Println(r.ContentLength)
}

func main() {
	log.Println("create router")
	router := httprouter.New()

	log.Println("register user")
	handler := user.NewHandler()
	handler.Register(router)

	log.Println("start router")
	start(router)
}

func start(router *httprouter.Router) {
	log.Println("start listener")

	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second}

	log.Println("server listening on port 8080")
	log.Fatal(server.Serve(listener))
}
