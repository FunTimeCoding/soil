package client

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/alpine/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/web"
	"path/filepath"
)

func (c *Client) Upload(
	path string,
	version string,
	repository string,
	architecture string,
) (string, int) {
	q := web.NewPostBytes(
		join.Slash(
			[]string{
				c.base,
				constant.RoutePrefix,
				version,
				repository,
				architecture,
				filepath.Base(path),
			},
		),
		bytes.NewReader(system.ReadBytesUnsafe(path)),
	)
	web.Bearer(q, c.token)
	r, e := web.Client().Do(q)
	errors.PanicOnError(e)
	defer errors.PanicClose(r.Body)

	return web.ReadString(r), r.StatusCode
}
