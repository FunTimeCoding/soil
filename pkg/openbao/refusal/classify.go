package refusal

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/openbao/openbao/api/v2"
	"net/http"
)

func Classify(e error) error {
	if e == nil {
		return nil
	}

	if f := connection.Classify(e); f != nil {
		return f
	}

	var answer *api.ResponseError

	if errors.As(e, &answer) &&
		answer.StatusCode == http.StatusServiceUnavailable {
		return unavailable.Format("%s", e.Error())
	}

	return e
}
