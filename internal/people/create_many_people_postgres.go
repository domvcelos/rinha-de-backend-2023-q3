package people

import (
	"context"
	"fmt"
	"strings"
)

func (c *PeoplePostgres) CreateMany(ctx context.Context, peoples []*People) error {
	arraySize := len(peoples)
	collumnQuantity := 5
	valueStrings := make([]string, 0, arraySize)
	valueArgs := make([]interface{}, 0, arraySize*collumnQuantity)
	i := 0
	for _, p := range peoples {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)", i*collumnQuantity+1, i*collumnQuantity+2, i*collumnQuantity+3, i*collumnQuantity+4, i*collumnQuantity+5))
		valueArgs = append(valueArgs, p.Id)
		valueArgs = append(valueArgs, p.Apelido)
		valueArgs = append(valueArgs, p.Nome)
		valueArgs = append(valueArgs, p.Nascimento)
		valueArgs = append(valueArgs, strings.Join(p.Stack[:], "-"))
		i++
	}
	stmt := fmt.Sprintf("INSERT INTO public.people (id, apelido, nome, nascimento, stack) VALUES %s", strings.Join(valueStrings, ","))
	_, err := c.DB.Exec(stmt, valueArgs...)
	return err
}
