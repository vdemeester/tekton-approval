package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vdemeester/tekton-approval/pkg/approval"
)

var (
	port = flag.Int("port", 8080, "port to expose the app on")
)

func main() {
	flag.Parse()
	p := fmt.Sprintf(":%v", *port)
	log.Println("Run approval on", p)

	r := mux.NewRouter()
	h := approval.NewHTTPHandler()

	r.HandleFunc("/approval", h.Get).Methods("GET")

	log.Fatal(http.ListenAndServe(p, r))
}
