package entries

import (
	"fmt"
	"log"
	"net/http"
)

type Handler struct{}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	log.Println("received request to create a entry")
	_, err := w.Write([]byte("Entry created!"))
	if err != nil {
		log.Printf("Failed to create entry %v", err)
	}
}

func (h *Handler) FindByID(w http.ResponseWriter, r *http.Request) {
	log.Println("handling READ request - Method:", r.Method)
	w.WriteHeader(http.StatusOK)
	res := "Coffee " + r.PathValue("id")
	fmt.Fprint(w, res)
	// monster, exists := loadMonsters()[r.PathValue("id")]
	// if !exists {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	return
	// }
	// json.NewEncoder(w).Encode(monster)
}

func (h *Handler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling UPDATE request - Method:", r.Method)
}

func (h *Handler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	log.Println("received DELETE request for monster")
}
