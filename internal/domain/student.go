package domain

import "errors"

type Student struct {
	ID    int
	Name  string
	Age   int
	Tel   string
	Major string
}

func NewStudent(id int, name string, age int, tel, major string) (*Student, error) {
	if name == "" || age <= 0 || tel == "" || major == "" {
		return nil, errors.New("invalid student data")
	}
	return &Student{
		ID:    id,
		Name:  name,
		Age:   age,
		Tel:   tel,
		Major: major,
	}, nil
}
