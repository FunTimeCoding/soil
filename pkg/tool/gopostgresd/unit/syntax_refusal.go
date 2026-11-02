package unit

import "github.com/funtimecoding/soil/pkg/errors/validation"

func syntaxRefusal() error {
	return validation.New("postgres: syntax error at or near \"SELEC\"")
}
