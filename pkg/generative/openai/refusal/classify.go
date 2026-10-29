package refusal

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"github.com/openai/openai-go/v3"
)

func Classify(e error) error {
	if e == nil {
		return nil
	}

	if f := connection.Classify(e); f != nil {
		return f
	}

	var s *openai.Error

	if !errors.As(e, &s) {
		return e
	}

	if classified := refusal.Status(
		s.StatusCode,
		join.Empty("openai: ", s.Message),
	); classified != nil {
		return classified
	}

	return unexpected.Format("openai status: %d: %s", s.StatusCode, s.Message)
}
