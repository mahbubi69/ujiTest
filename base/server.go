package base

import (
	"encoding/json"
	"net/http"
	"ujiTest/res"
)

type Server struct {
	Port string
}

func (s *Server) StatusServer(w http.ResponseWriter, r *http.Request) {

	response := res.Status{
		Status:  200,
		Message: "Server Is Acctived!!",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
