package people

import (
	"encoding/json"
	"net/http"

	"github.com/domvcelos/rinha-de-backend-2023-q3/pkg"
)

func (handler *PeopleHandler) Count(w http.ResponseWriter, r *http.Request) {
	resp, err := handler.Service.Count(r.Context())
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := pkg.Error{Message: err.Error()}
		_ = json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
