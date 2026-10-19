package basic

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func rejected(
	status int,
	body []byte,
) error {
	if status != http.StatusBadRequest {
		return nil
	}

	return validation.New("%s", refusal.Reason(body, http.StatusText(status)))
}
