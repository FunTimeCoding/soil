package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/backrest/response"
)

type BackrestClient interface {
	Stanzas(x context.Context) ([]*response.Stanza, error)
	Reachable(x context.Context) error
}
