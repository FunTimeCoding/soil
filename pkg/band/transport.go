package band

import "github.com/funtimecoding/soil/pkg/errors/connection"

func transport(e error) error {
	if f := connection.Classify(e); f != nil {
		return f
	}

	return e
}
