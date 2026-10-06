package classify

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/errors/classify/face"
)

func expected(e error) bool {
	var x face.Expected

	return errors.As(e, &x) && x.Expected()
}
