package base

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"ujiTest/models"
	"ujiTest/res"

	"github.com/gorilla/mux"
)

var data []models.Item

var modelReferences = []string{"car", "humanoid", "transformation"}
var techReferences = []string{"AI", "car", "robot", "cyborg", "cybord"}

func (s *Server) GetItems(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		modelFilter := r.URL.Query().Get("model")
		techFilters := r.URL.Query()["tech"]

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
		response := res.GetResponse{
			Status:     http.StatusOK,
			Count:      len(filteredItems),
			TotalCount: len(data),
			Data:       filteredItems,
		}

		// Mengirimkan response
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

		// Add new item to data
		data = append(data, newItem)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newItem)
	}
}

// func (s *Server) GetItems(w http.ResponseWriter, r *http.Request) {
// 	filePath := "data.txt"
// 	data := s.LoadDataFromFile(filePath)

// 	switch r.Method {
// 	case "GET":
// 		// Ambil filter dari URL query
// 		modelFilter := r.URL.Query().Get("model")
// 		techFilters := r.URL.Query()["tech"]

// 		var filteredItems []models.Item
// 		for _, item := range data {
// 			// Filter berdasarkan model jika diberikan
// 			if modelFilter != "" && item.Model != modelFilter {
// 				continue
// 			}
// 			// Filter berdasarkan teknologi jika ada
// 			if len(techFilters) > 0 {
// 				matches := true
// 				for _, tech := range techFilters {
// 					if !s.Contains(item.Tech, tech) {
// 						matches = false
// 						break
// 					}
// 				}
// 				if !matches {
// 					continue
// 				}
// 			}
// 			filteredItems = append(filteredItems, item)
// 		}

// 		// Siapkan response
// 		response := res.GetResponse{
// 			Status:     http.StatusOK,
// 			Count:      len(filteredItems),
// 			TotalCount: len(data),
// 			Data:       filteredItems,
// 		}

// 		// Kirimkan response
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusOK)
// 		json.NewEncoder(w).Encode(response)

// 	case "POST":
// 		// Tambahkan item baru
// 		var newItem models.Item
// 		err := json.NewDecoder(r.Body).Decode(&newItem)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}

// 		// Tambahkan ke data (untuk sementara tidak disimpan ke file)
// 		data = append(data, newItem)

// 		// Kirimkan response
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusCreated)
// 		json.NewEncoder(w).Encode(newItem)
// 	default:
// 		// Tangani metode yang tidak didukung
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 	}
// }

// func (s *Server) GetItems(w http.ResponseWriter, r *http.Request) {
// 	filePath := "data.txt"

// 	// Gunakan channel untuk sinkronisasi data dari goroutine
// 	dataChan := make(chan []models.Item)
// 	errChan := make(chan error)

// 	// Jalankan `LoadDataFromFile` dalam goroutine
// 	go func() {
// 		data, err := s.LoadDataFromFile(filePath)
// 		if err != nil {
// 			errChan <- err
// 			return
// 		}
// 		dataChan <- data
// 	}()

// 	var data []models.Item
// 	select {
// 	case data = <-dataChan: // Berhasil memuat data
// 	case err := <-errChan: // Terjadi error
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	switch r.Method {
// 	case "GET":
// 		// Ambil filter dari URL query
// 		modelFilter := r.URL.Query().Get("model")
// 		techFilters := r.URL.Query()["tech"]

// 		filteredChan := make(chan []models.Item)

// 		// Jalankan proses filter dalam goroutine
// 		go func() {
// 			var filteredItems []models.Item
// 			for _, item := range data {
// 				// Filter berdasarkan model jika diberikan
// 				if modelFilter != "" && item.Model != modelFilter {
// 					continue
// 				}
// 				// Filter berdasarkan teknologi jika ada
// 				if len(techFilters) > 0 {
// 					matches := true
// 					for _, tech := range techFilters {
// 						if !s.Contains(item.Tech, tech) {
// 							matches = false
// 							break
// 						}
// 					}
// 					if !matches {
// 						continue
// 					}
// 				}
// 				filteredItems = append(filteredItems, item)
// 			}
// 			filteredChan <- filteredItems
// 		}()

// 		// Tunggu hasil filter
// 		filteredItems := <-filteredChan

// 		// Siapkan response
// 		response := res.GetResponse{
// 			Status:     http.StatusOK,
// 			Count:      len(filteredItems),
// 			TotalCount: len(data),
// 			Data:       filteredItems,
// 		}

// 		// Kirimkan response
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusOK)
// 		json.NewEncoder(w).Encode(response)

// 	case "POST":
// 		// Tambahkan item baru
// 		var newItem models.Item
// 		err := json.NewDecoder(r.Body).Decode(&newItem)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}

// 		// Tambahkan ke data (untuk sementara tidak disimpan ke file)
// 		go func() {
// 			data = append(data, newItem) // Tambahkan dalam goroutine
// 		}()

// 		// Kirimkan response
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusCreated)
// 		json.NewEncoder(w).Encode(newItem)

// 	default:
// 		// Tangani metode yang tidak didukung
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 	}
// }

func (s *Server) GetItemByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	// Cari item berdasarkan code
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

	// Debug log untuk melihat nilai code
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

func (s *Server) GetReferences(w http.ResponseWriter, r *http.Request) {
	refType := strings.TrimPrefix(r.URL.Path, "/references/")

	var references []string
	if refType == "model" {
		references = modelReferences
	} else if refType == "tech" {
		references = techReferences
	} else {
		http.Error(w, "Invalid reference type", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(references)
}

func (s *Server) Contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
