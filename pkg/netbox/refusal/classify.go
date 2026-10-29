package refusal

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/netbox-community/go-netbox/v4"
)

func Classify(e error) error {
	if e == nil {
		return nil
	}

	if f := connection.Classify(e); f != nil {
		return f
	}

	var g *netbox.GenericOpenAPIError

	if !errors.As(e, &g) {
		return e
	}

	return New(status(g.Error()), reason(g.Body(), g.Error()))
}
