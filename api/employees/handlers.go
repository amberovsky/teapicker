package employees

import (
	"amberovsky/teapicker/internal"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"sort"
)

/*
Handlers for /employee endpoint - employees management
*/

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		return
	}

	employees, _ := h.service.GetAll()
	sort.Slice(employees, func(i, j int) bool {
		return employees[i].Name < employees[j].Name
	})

	if err := json.NewEncoder(w).Encode(employees); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

type addEmployeeRequest struct {
	Name string `json:"name"`
}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var ae addEmployeeRequest
	_ = json.NewDecoder(r.Body).Decode(&ae)

	internal.HandleHttpErr(w, h.service.Add(ae.Name))
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == http.MethodOptions {
		return
	}

	vars := mux.Vars(r)

	internal.HandleHttpErr(w, h.service.Delete(vars["uuid"]))
}
