package base

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) Routes() {

	s.SeedItemHelper()

	r := mux.NewRouter()

	r.HandleFunc("/status", s.StatusServer).Methods("GET")
	r.HandleFunc("/items", s.GetItems).Methods("GET", "POST")           // GET and POST to manage items
	r.HandleFunc("/items/{code}", s.GetItemByCode).Methods("GET")       // GET by code
	r.HandleFunc("/updateItems/{code}", s.UpdateItem).Methods("PUT")    //PUT by code
	r.HandleFunc("/deleteItems/{code}", s.DeleteItem).Methods("DELETE") // DELETE by code
	r.HandleFunc("/references/", s.GetReferences).Methods("GET")

	// Menjalankan HTTP server
	http.Handle("/", r)
	http.ListenAndServe(":"+s.Port, nil)
}

func NewServer(port string) *Server {
	return &Server{Port: port}
}
