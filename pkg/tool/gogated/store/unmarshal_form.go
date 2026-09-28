package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"net/url"
)

func unmarshalForm(s string) url.Values {
	v, e := url.ParseQuery(s)
	errors.PanicOnError(e)

	return v
}
