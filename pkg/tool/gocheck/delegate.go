package gocheck

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/option"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/types/check_result"
	gosourced "github.com/funtimecoding/soil/pkg/tool/gosourced/client"
	gosourcedConstant "github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/generated/client"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func Delegate(o *option.Check) (*check_result.Result, bool, error) {
	token := environment.Optional(gosourcedConstant.TokenEnvironment)

	if token == "" {
		return nil, false, nil
	}

	c := gosourced.New(
		locator.Environment(
			gosourcedConstant.HostEnvironment,
			gosourcedConstant.PortEnvironment,
			gosourcedConstant.InsecureEnvironment,
		),
		token,
	)
	defer c.Close()
	x := context.Background()

	if c.Reachable(x) != nil {
		return nil, false, nil
	}

	r, e := c.PostCheck(
		x,
		client.CheckRequest{
			Root:       o.Root,
			Fix:        new(o.Fix),
			Fresh:      new(o.Fresh),
			Transcript: new(o.Transcript),
		},
	)

	if not_found.Is(e) {
		return nil, false, nil
	}

	if e != nil {
		return nil, false, e
	}

	return fromResponse(r), true, nil
}
