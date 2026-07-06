package brew

import (
	"github.com/funtimecoding/soil/pkg/system/macos/brew/constant"
	"github.com/funtimecoding/soil/pkg/system/macos/brew/response"
	"github.com/funtimecoding/soil/pkg/system/run"
)

func (c *Client) Installed() *response.Installed {
	r := run.New()
	r.Start(
		constant.Brew,
		constant.Information,
		constant.Installed,
		constant.Notation2,
	)
	r.Verbose = true
	var result response.Installed
	r.ParseNotation(&result)

	return &result
}
