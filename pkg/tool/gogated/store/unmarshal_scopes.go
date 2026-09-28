package store

import (
	"github.com/ory/fosite"
	"strings"
)

func unmarshalScopes(s string) fosite.Arguments {
	if s == "" {
		return fosite.Arguments{}
	}

	return strings.Split(s, " ")
}
