package refusal

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	errorsConstant "github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"github.com/ollama/ollama/api"
	"net/http"
)

func Classify(e error) error {
	if e == nil {
		return nil
	}

	if f := connection.Classify(e); f != nil {
		return f
	}

	var s api.StatusError

	if !errors.As(e, &s) {
		return e
	}

	if classified := refusal.NewTable().With(
		http.StatusBadRequest,
		errorsConstant.ValidationClass,
	).Classify(
		s.StatusCode,
		join.Empty("ollama: ", s.Error()),
	); classified != nil {
		return classified
	}

	return unexpected.Format("ollama status: %d: %s", s.StatusCode, s.Error())
}
