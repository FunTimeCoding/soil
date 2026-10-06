package client

import (
	"github.com/funtimecoding/soil/pkg/nextcloud/constant"
	"github.com/funtimecoding/soil/pkg/nextcloud/usage"
	"github.com/funtimecoding/soil/pkg/nextcloud/usage/response"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
)

func (c *Client) Fetch() (*usage.Usage, error) {
	p := response.NewPayload()

	if e := c.requester.Notation(
		request.Get(constant.InformationPath).WithParameter(
			constant.FormatParameter,
			constant.NotationFormat,
		),
		p,
	); e != nil {
		return nil, e
	}

	return usage.New(
		p.Wrapper.Body.Nextcloud.Storage.Files,
		p.Wrapper.Body.Nextcloud.Shares.Count,
	), nil
}
