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
	port        = flag.Int("port", 8080, "port to expose the app on")
	triggersURL = flag.String("triggers", "", "URL to the triggers service")
)

func main() {
	flag.Parse()
	p := fmt.Sprintf(":%v", *port)
	log.Println("Run approval on", p)

	r := mux.NewRouter()
	h := approval.NewHTTPHandler()

	r.HandleFunc("/approval", h.List).Methods("GET")
	r.HandleFunc("/approval", h.Add).Methods("POST")
	r.HandleFunc("/approval/{id}", h.Get).Methods("GET")
	r.HandleFunc("/approval/{id}", h.Update(*triggersURL)).Methods("PUT")

	log.Fatal(http.ListenAndServe(p, r))
}
