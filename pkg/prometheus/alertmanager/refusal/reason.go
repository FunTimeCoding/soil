package refusal

import (
	"errors"
	"net/http"
)

func reason(
	e error,
	code int,
) string {
	if c, okay := errors.AsType[carrying](e); okay && c.GetPayload() != "" {
		return c.GetPayload()
	}

	return http.StatusText(code)
}
