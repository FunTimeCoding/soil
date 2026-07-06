package debian

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
)

func New() *Client {
	h := system.Home()

	return &Client{
		home:          h,
		workDirectory: join.Absolute(h, constant.DownloadsPath),
	}
}
