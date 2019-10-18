package approval

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct{}

func NewHTTPHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Get(w http.ResponseWriter, req *http.Request) {
	log.Println("GET /")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Get()); err != nil {
		// TODO: return an error
		fmt.Println("error getting approvals", err)
	}
}

func (h *Handler) Add(w http.ResponseWriter, req *http.Request) {
	var approval Approval
	if err := json.NewDecoder(req.Body).Decode(&approval); err != nil {
		// TODO: return an error
		fmt.Println("error creating approvals", err)
	}
	log.Printf("POST with %v", approval)
	if approval.Status == "" {
		approval.Status = StatusUnknown
	}
	created := Add(approval.Name, approval.URL, approval.Status)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(created); err != nil {
		// TODO: return an error
		fmt.Println("error getting approvals", err)
	}
}

func (h *Handler) Update(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var approval Approval
	if err := json.NewDecoder(req.Body).Decode(&approval); err != nil {
		// TODO: return an error
		fmt.Println("error creating approvals", err)
	}
	log.Printf("PUT with %v", approval)
	if approval.ID == "" {
		approval.ID = params["id"]
	}
	if approval.Status == "" {
		approval.Status = StatusUnknown
	}
	created := Update(approval.ID, approval)
	w.WriteHeader(http.StatusAccepted)
	if err := json.NewEncoder(w).Encode(created); err != nil {
		// TODO: return an error
		fmt.Println("error getting approvals", err)
	}
}
