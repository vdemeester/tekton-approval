package approval

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type Handler struct{}

func NewHTTPHandler() *Handler {
	return &Handler{}
}

func (h *Handler) List(w http.ResponseWriter, req *http.Request) {
	log.Println("GET /")
	CORSEnabledFunction(w, req)
	if err := json.NewEncoder(w).Encode(List()); err != nil {
		// TODO: return an error
		fmt.Println("error getting approvals", err)
	}
}
func (h *Handler) Options(w http.ResponseWriter, req *http.Request) {
	log.Println("OPTIONS /")
	CORSEnabledFunction(w, req)
}

func (h *Handler) Add(w http.ResponseWriter, req *http.Request) {
	var approval Approval
	if err := json.NewDecoder(req.Body).Decode(&approval); err != nil {
		// TODO: return an error
		fmt.Println("error creating approvals", err)
	}
	CORSEnabledFunction(w, req)
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

type event struct {
	from string
	to   string
}

func (h *Handler) Update(url string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		var approval Approval
		if err := json.NewDecoder(req.Body).Decode(&approval); err != nil {
			// TODO: return an error
			fmt.Println("error creating approvals", err)
		}
		CORSEnabledFunction(w, req)
		log.Printf("PUT with %v", approval)
		if approval.ID == "" {
			approval.ID = params["id"]
		}
		if approval.Status == "" {
			approval.Status = StatusUnknown
		}
		if approval.Status == StatusApproved && url != "" { // HACK HACK HACK
			e := event{from: "dev", to: "to"}
			data, err := json.Marshal(e)
			if err != nil {
				log.Println("Error marshaling the error. ", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			req, err := http.NewRequest("POST", url, strings.NewReader(string(data)))
			if err != nil {
				log.Println("Error creating the request. ", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			req.Header.Set("Content-Type", "application/json")
			client := &http.Client{Timeout: time.Second * 10}

			resp, err := client.Do(req)
			if err != nil {
				log.Println("Error reading response. ", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println("Error reading body. ", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			log.Printf("%s\n", body)
		}
		created := Update(approval.ID, approval)
		w.WriteHeader(http.StatusAccepted)
		if err := json.NewEncoder(w).Encode(created); err != nil {
			// TODO: return an error
			fmt.Println("error getting approvals", err)
		}
	}
}

func (h *Handler) Get(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	CORSEnabledFunction(w, req)
	log.Printf("GET with /%v", id)
	a := Get(id)
	if a != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(a); err != nil {
			// TODO: return an error
			fmt.Println("error getting approvals", err)
		}
	}
}

func CORSEnabledFunction(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST,GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusOK)
		return
	}
	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
