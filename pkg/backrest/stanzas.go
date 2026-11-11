package backrest

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/backrest/constant"
	"github.com/funtimecoding/soil/pkg/backrest/response"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/notation"
)

func (c *Client) Stanzas(x context.Context) ([]*response.Stanza, error) {
	output, e := c.run(x, constant.InformationCommand, constant.NotationOutput)

	if e != nil {
		if f := c.Reachable(x); f != nil {
			return nil, unavailable.Format(constant.CommandUnavailable, f)
		}

		return nil, e
	}

	var result []*response.Stanza

	if f := notation.Decode(output, &result); f != nil {
		return nil, fmt.Errorf("decode stanzas: %w", f)
	}

	return result, nil
}
