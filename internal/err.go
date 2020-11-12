package internal

import (
	"errors"
	"net/http"
)

var ErrNotFound = errors.New("not found")

func HandleHttpErr(w http.ResponseWriter, err error) {
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
