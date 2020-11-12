package teams

import (
	"amberovsky/teapicker/internal"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"sort"
)

/*
Handlers for /team endpoint - managing teams
*/

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetTeams(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		return
	}

	teams, _ := h.service.GetAll()
	sort.Slice(teams, func(i, j int) bool {
		return teams[i].Name < teams[j].Name
	})

	internal.HandleHttpErr(w, json.NewEncoder(w).Encode(teams))
}

type addTeamRequest struct {
	Name string `json:"name"`
}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var at addTeamRequest
	_ = json.NewDecoder(r.Body).Decode(&at)

	internal.HandleHttpErr(w, h.service.Add(at.Name))
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
