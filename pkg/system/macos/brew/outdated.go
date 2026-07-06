package brew

import (
	"github.com/funtimecoding/soil/pkg/system/macos/brew/constant"
	"github.com/funtimecoding/soil/pkg/system/macos/brew/response"
	"github.com/funtimecoding/soil/pkg/system/run"
)

func (c *Client) Outdated() *response.Outdated {
	r := run.New()
	r.Start(constant.Brew, constant.Outdated, constant.Notation2)
	var result response.Outdated
	r.ParseNotation(&result)

	return &result
}
