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
	// http://el-deploy-listener.demo.svc.cluster.local
	triggersURL = flag.String("triggers", "", "URL to the triggers service")
)

func main() {
	flag.Parse()
	p := fmt.Sprintf(":%v", *port)
	log.Println("Run approval on", p)

	r := mux.NewRouter()
	h := approval.NewHTTPHandler()

	// Create room for static files serving
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("dist/tekton-approval//"))).Methods("GET")
	r.HandleFunc("/approval", h.Options).Methods("OPTIONS")
	r.HandleFunc("/approval", h.List).Methods("GET")
	r.HandleFunc("/approval", h.Add).Methods("POST")
	r.HandleFunc("/approval/{id}", h.Get).Methods("GET")
	r.HandleFunc("/approval/{id}", h.Update(*triggersURL)).Methods("PUT")

	log.Fatal(http.ListenAndServe(p, r))
}
