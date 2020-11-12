package members

import (
	"amberovsky/teapicker/internal"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

/*
Handlers for /member endpoint - managing teams composition (association)
*/

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetMembers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		return
	}

	teamUUID := r.URL.Query().Get("team_uuid")

	internal.HandleHttpErr(
		w,
		json.NewEncoder(w).Encode(h.service.GetTeamMembers(uuid.MustParse(teamUUID))),
	)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		return
	}

	vars := mux.Vars(r)

	internal.HandleHttpErr(w, json.NewEncoder(w).Encode(h.service.Get(uuid.MustParse(vars["uuid"]))))
}

type addMemberRequest struct {
	TeamUUID     string `json:"team_uuid"`
	EmployeeUUID string `json:"employee_uuid"`
}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var am addMemberRequest
	_ = json.NewDecoder(r.Body).Decode(&am)

	internal.HandleHttpErr(w, h.service.Add(am.TeamUUID, am.EmployeeUUID))
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)

	internal.HandleHttpErr(w, h.service.Delete(vars["uuid"]))
}
