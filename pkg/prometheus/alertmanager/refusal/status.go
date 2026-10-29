package refusal

import (
	"errors"
	"github.com/go-openapi/runtime"
)

func status(e error) (int, bool) {
	if c, okay := errors.AsType[coded](e); okay {
		return c.Code(), true
	}

	if a, okay := errors.AsType[*runtime.APIError](e); okay {
		return a.Code, true
	}

	return 0, false
}
