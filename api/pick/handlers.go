package pick

import (
	"amberovsky/teapicker/internal"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

/*
Handlers for /pick endpoint - picking the next team maker
*/

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

type pickResponse struct {
	MemberUUID uuid.UUID `json:"member_uuid"`
}

func (h *Handler) Pick(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == http.MethodOptions {
		return
	}

	vars := mux.Vars(r)
	memberUUUID, _ := h.service.Pick(vars["uuid"])

	if memberUUUID != nil {
		err := json.NewEncoder(w).Encode(&pickResponse{
			MemberUUID: *memberUUUID,
		})

		internal.HandleHttpErr(w, err)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
