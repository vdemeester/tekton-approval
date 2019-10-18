package approval

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Handler struct{}

func NewHTTPHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Get(w http.ResponseWriter, req *http.Request) {
	if err := json.NewEncoder(w).Encode(Get()); err != nil {
		// TODO: return an error
		fmt.Println("error getting approvals", err)
	}
}
