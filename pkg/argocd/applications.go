package argocd

import (
	"github.com/funtimecoding/soil/pkg/argocd/application"
	"github.com/funtimecoding/soil/pkg/argocd/constant"
	"github.com/funtimecoding/soil/pkg/argocd/response"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
)

func (c *Client) Applications() ([]*application.Application, error) {
	var p response.Applications

	if e := c.requester.Notation(request.Get(constant.ApplicationsPath), &p); e != nil {
		return nil, e
	}

	return application.NewSlice(p.Items), nil
}
