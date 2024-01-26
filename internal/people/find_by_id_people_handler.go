package people

import (
	"encoding/json"
	"net/http"

	"github.com/domvcelos/rinha-de-backend-2023-q3/pkg"
	"github.com/go-chi/chi/v5"
)

func (handler *PeopleHandler) FindById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "peopleID")
	resp, err := handler.Service.FindById(r.Context(), id)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := pkg.Error{Message: err.Error()}
		_ = json.NewEncoder(w).Encode(error)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
