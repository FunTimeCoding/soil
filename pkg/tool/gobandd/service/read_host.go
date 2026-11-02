package service

import (
	"github.com/funtimecoding/soil/pkg/band/face"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/types/status"
	"strings"
)

func readHost(
	c face.BandClient,
	v *status.Status,
) error {
	r, e := c.GeneralSettings()

	if e != nil {
		return e
	}

	v.Host = strings.Trim(
		join.Dot([]string{r.HostName, r.DomainName}),
		constant.Dot,
	)

	return nil
}
