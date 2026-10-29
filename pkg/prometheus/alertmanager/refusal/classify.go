package refusal

import (
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func Classify(e error) error {
	if e == nil {
		return nil
	}

	if f := connection.Classify(e); f != nil {
		return f
	}

	code, okay := status(e)

	if !okay {
		return e
	}

	text := reason(e, code)

	if classified := refusal.Status(
		code,
		join.Empty("alertmanager: ", text),
	); classified != nil {
		return classified
	}

	return unexpected.Format("alertmanager status: %d: %s", code, text)
}
