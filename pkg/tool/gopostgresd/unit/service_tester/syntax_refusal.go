package service_tester

import "github.com/funtimecoding/soil/pkg/errors/validation"

func SyntaxRefusal() error {
	return validation.New("postgres: syntax error at or near \"SELEC\"")
}
