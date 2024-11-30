package base

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"ujiTest/models"
	"ujiTest/res"

	"github.com/gorilla/mux"
)

var data []models.Item

// var modelReferences = []string{"car", "humanoid", "transformation"}
// var techReferences = []string{"AI", "car", "robot", "cyborg", "cybord"}

func (s *Server) GetItems(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		// get query parameter URL
		modelFilter := r.URL.Query().Get("model")
		techFilterQuery := r.URL.Query().Get("tech")

		var techFilters []string
		if techFilterQuery != "" {
			techFilters = strings.Split(techFilterQuery, ",")
		}

		var filteredItems []models.Item
		for _, item := range data {
			if modelFilter != "" && item.Model != modelFilter {
				continue
			}

			if len(techFilters) > 0 {
				matches := true
				for _, tech := range techFilters {
					if !s.Contains(item.Tech, tech) {
						matches = false
						break
					}
				}
				if !matches {
					continue
				}
			}
			filteredItems = append(filteredItems, item)
		}

		if filteredItems == nil {
			filteredItems = []models.Item{}
		}

		response := res.GetResponse{
			Status:     http.StatusOK,
			Count:      len(filteredItems),
			TotalCount: len(data),
			Data:       filteredItems,
		}

		// send response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)

	case "POST":
		var newItem models.Item
		err := json.NewDecoder(r.Body).Decode(&newItem)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if len(newItem.Code) < 10 {
			json.NewEncoder(w).Encode(res.Status{Status: 400, Message: "Maaf input data anda kurang dari 10 "})
			return
		}

		if len(newItem.Code) > 15 {
			json.NewEncoder(w).Encode(res.Status{Status: 400, Message: "Maaf input data tidak boleh lebih dari 15"})
			return
		}

		validate, _ := regexp.MatchString(`^[a-zA]+$`, newItem.Code)
		if !validate {
			json.NewEncoder(w).Encode(res.Status{Status: 400, Message: "Maaf input data hanya berisi angka dan huruf"})
			return
		}

		for _, item := range data {
			if item.Code == newItem.Code {
				json.NewEncoder(w).Encode(res.Status{Status: 400, Message: "Data dengan code sudah ada"})
				return
			}
			if item.Model == newItem.Model {
				json.NewEncoder(w).Encode(res.Status{Status: 400, Message: "Data dengan model sudah ada"})
				return
			}
		}

		data = append(data, newItem)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newItem)
	}
}

func (s *Server) GetItemByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	//  search by code
	for _, item := range data {
		if item.Code == code {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}

func (s *Server) UpdateItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	fmt.Println("Received code:", code)

	if code == "" {
		http.Error(w, "Code parameter is missing", http.StatusBadRequest)
		return
	}

	var updatedItem models.Item
	err := json.NewDecoder(r.Body).Decode(&updatedItem)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	for i, item := range data {
		if item.Code == code {

			data[i] = updatedItem

			data[i].Code = code

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Item updated successfully"))
			return
		}
	}

	http.Error(w, "Item not found", http.StatusNotFound)
}

func (s *Server) DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	// Debug log
	fmt.Println("Received code:", code)

	if code == "" {
		http.Error(w, "Code parameter is missing", http.StatusBadRequest)
		return
	}

	for i, item := range data {
		if item.Code == code {
			data = append(data[:i], data[i+1:]...)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Item deleted successfully"))
			return
		}
	}

	http.Error(w, "Item not found", http.StatusNotFound)
}

// func (s *Server) GetReferences(w http.ResponseWriter, r *http.Request) {
// 	refType := strings.TrimPrefix(r.URL.Path, "/references/")

// 	var references []string
// 	if refType == "model" {
// 		references = modelReferences
// 	} else if refType == "tech" {
// 		references = techReferences
// 	} else {
// 		http.Error(w, "Invalid reference type", http.StatusBadRequest)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(references)
// }

func (s *Server) Contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// helper send to memory data json
func (s *Server) SeedItemHelper() {
	var items []models.Item

	// open files json
	jsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successfully opened data.json")
	defer jsonFile.Close()

	// read files json dari file JSON
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Unmarshal JSON to struct Item
	err = json.Unmarshal(byteValue, &items)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Proses save data global
	for _, item := range items {
		data = append(data, item)
		fmt.Printf("Item saved: %v\n", item.Name)
	}

	// result
	fmt.Println("All items saved in memory")
}
