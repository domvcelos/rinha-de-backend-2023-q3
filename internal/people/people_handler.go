package people

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/domvcelos/rinha-de-backend-2023-q3/pkg"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type PeopleHandler struct {
	Service PeopleServiceInterface
}

func NewHandler(ps PeopleServiceInterface) *PeopleHandler {
	return &PeopleHandler{
		Service: ps,
	}
}

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

func (handler *PeopleHandler) FindById(w http.ResponseWriter, r *http.Request) {
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
func (handler *PeopleHandler) Find(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("t")
	w.Header().Set("Content-Type", "application/json")
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

func (handler *PeopleHandler) Count(w http.ResponseWriter, r *http.Request) {
	resp, err := handler.Service.Count(r.Context())
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
