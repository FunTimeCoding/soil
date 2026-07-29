package brew

import (
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/macos/brew/response"
	"github.com/funtimecoding/soil/pkg/system/run"
)

func (c *Client) Outdated() *response.Outdated {
	r := run.New()
	r.Start(constant.Brew, constant.BrewOutdated, constant.BrewNotation2)
	var result response.Outdated
	r.ParseNotation(&result)

	return &result
}
