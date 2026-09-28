package store

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/ory/fosite"
	"github.com/ory/fosite/handler/openid"
)

func unmarshalSession(
	s string,
	e fosite.Session,
) fosite.Session {
	if s == "" {
		return e
	}

	if e == nil {
		e = &openid.DefaultSession{}
	}

	errors.PanicOnError(json.Unmarshal([]byte(s), e))

	return e
}
