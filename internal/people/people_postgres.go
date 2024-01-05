package people

import (
	"context"
	"database/sql"
	"strings"
)

type PeoplePosgres struct {
	DB *sql.DB
}

func NewPostgres(db *sql.DB) *PeoplePosgres {
	return &PeoplePosgres{DB: db}
}

func (client *PeoplePosgres) Create(ctx context.Context, people *People) error {
	stmt, err := client.DB.Prepare(`INSERT INTO public.people
	(id, apelido, nome, nascimento, stack)
	VALUES(gen_random_uuid(), $1, $2, $3, $4);`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(people.Apelido, people.Nome, people.Nascimento, strings.Join(people.Stack[:], " "))
	if err != nil {
		return err
	}
	return nil
}
func (client *PeoplePosgres) FindById(ctx context.Context, id string) (*People, error) {
	stmt, err := client.DB.Prepare(`select p.apelido, p.id, p.nome, p.nascimento, p.stack from public.people p where p.id = $1;`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	p := People{}
	var stacks string
	err = stmt.QueryRow(id).Scan(&p.Apelido, &p.Id, &p.Nome, &p.Nascimento, &stacks)
	if err != nil {
		return nil, err
	}
	if stacks != "" {
		p.Stack = strings.Split(stacks, " ")
	}
	return &p, nil
}
func (client *PeoplePosgres) Find(ctx context.Context, query string) (*[]People, error) {
	stmt, err := client.DB.Prepare(`select p.apelido, p.id, p.nome, p.nascimento, p.stack from public.people p where p.busca_tgrm ilike '%'||$1||'%' LIMIT 50;`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	peoples := []People{}

	var rows *sql.Rows
	if query == "" {
		rows, err = client.DB.Query(`select p.apelido, p.id, p.nome, p.nascimento, p.stack from public.people p LIMIT 50`)
	} else {
		rows, err = stmt.Query(query)
	}
	for rows.Next() {
		p := People{}
		var stacks string
		err = rows.Scan(&p.Apelido, &p.Id, &p.Nome, &p.Nascimento, &stacks)
		if stacks != "" {
			p.Stack = strings.Split(stacks, " ")
		}
		peoples = append(peoples, p)
	}
	if err != nil {
		return nil, err
	}

	return &peoples, nil
}

func (client *PeoplePosgres) Count(ctx context.Context) (int, error) {
	var count int
	err := client.DB.QueryRow("select count(*) from public.people").Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
