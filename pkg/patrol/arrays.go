package patrol

import (
	"context"
	"github.com/funtimecoding/soil/pkg/mdstat"
	"github.com/funtimecoding/soil/pkg/mdstat/array"
	mdstatConstant "github.com/funtimecoding/soil/pkg/mdstat/constant"
)

func (c *Client) Arrays(_ context.Context) ([]*array.Array, error) {
	s, e := c.dial()

	if e != nil {
		return nil, e
	}

	defer s.Close()
	r := s.Run(mdstatConstant.ReadCommand)

	if r.Error != nil {
		return nil, r.Error
	}

	return mdstat.Parse(r.OutputString), nil
}
