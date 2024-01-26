package people

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/domvcelos/rinha-de-backend-2023-q3/pkg"
	"github.com/go-playground/validator/v10"
)

func (handler *PeopleHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto CreatePeopleDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := pkg.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(&dto)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		error := pkg.Error{Message: err.Error()}
		_ = json.NewEncoder(w).Encode(error)
		return
	}
	p, err := NewPeople(dto.Apelido, dto.Nome, dto.Nascimento, dto.Stack)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := pkg.Error{Message: err.Error()}
		_ = json.NewEncoder(w).Encode(error)
		return
	}
	id, err := handler.Service.Create(r.Context(), p)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		error := pkg.Error{Message: err.Error()}
		_ = json.NewEncoder(w).Encode(error)
		return
	}
	w.Header().Set("Location", "/pessoas/"+id)
	w.WriteHeader(http.StatusCreated)
}
