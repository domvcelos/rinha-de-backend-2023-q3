package people

import (
	"context"
	"fmt"
	"strings"
)

func (client *PeoplePostgres) Create(ctx context.Context, ch <-chan *People) {
	stmt, err := client.DB.Prepare(`INSERT INTO public.people
	(id, apelido, nome, nascimento, stack)
	VALUES($1, $2, $3, $4, $5) RETURNING id;`)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	for p := range ch {
		_, err = stmt.Exec(p.Id, p.Apelido, p.Nome, p.Nascimento, strings.Join(p.Stack[:], " "))
		if err != nil {
			fmt.Println(err)
		}
	}
	// err = stmt.QueryRow(p.Id, p.Apelido, p.Nome, p.Nascimento, strings.Join(p.Stack[:], " ")).Scan(&id)
	// if err != nil {
	// 	return err
	// }
}
