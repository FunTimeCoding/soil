package store

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/ory/fosite"
)

func marshalSession(s fosite.Session) string {
	b, e := json.Marshal(s)
	errors.PanicOnError(e)

	return string(b)
}
