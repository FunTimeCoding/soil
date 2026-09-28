package store

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/ory/fosite"
)

func marshalScopes(s fosite.Arguments) string {
	return join.Space(s...)
}
