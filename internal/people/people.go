package people

import (
	"fmt"
	"time"
)

type People struct {
	Id         string    `json:"id"`
	Apelido    string    `json:"apelido"`
	Nome       string    `json:"nome"`
	Nascimento time.Time `json:"nascimento"`
	Stack      []string  `json:"stack"`
}

type CreatePeopleDto struct {
	Apelido    string   `json:"apelido" validate:"required,max=32,min=1"`
	Nome       string   `json:"nome" validate:"max=100,min=1"`
	Nascimento string   `json:"nascimento" validate:"required"`
	Stack      []string `json:"stack" validate:"max=32"`
}

func NewPeople(apelido, nome, nascimento string, stack []string) (*People, error) {
	dates, err := ValidateDates(map[string]string{"nascimento": nascimento})
	if err != nil {
		return nil, err
	}
	if stack == nil {
		stack = nil
	}
	return &People{
		Apelido:    apelido,
		Nome:       nome,
		Nascimento: dates["nascimento"],
		Stack:      stack,
	}, nil
}
func ValidateDates(dates map[string]string) (map[string]time.Time, error) {
	parsedDates := make(map[string]time.Time)
	dateFormat := "2006-01-02"
	for key, date := range dates {
		parsedDate, err := time.Parse(dateFormat, date)
		if err != nil {
			return nil, fmt.Errorf("the field: %v is not in the correct format. Ex: %v", key, dateFormat)
		}
		parsedDates[key] = parsedDate
	}

	return parsedDates, nil
}
