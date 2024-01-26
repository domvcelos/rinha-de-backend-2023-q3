package people

import (
	"encoding/json"
	"net/http"

	"github.com/domvcelos/rinha-de-backend-2023-q3/pkg"
)

func (handler *PeopleHandler) Find(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query().Get("t")
	if query == "" {
		w.WriteHeader(http.StatusBadRequest)
		error := pkg.Error{Message: "Query string nao informada!"}
		json.NewEncoder(w).Encode(error)
		return
	}
	resp, err := handler.Service.Find(r.Context(), query)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := pkg.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
