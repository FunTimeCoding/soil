package argocd

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/argocd/application"
	"github.com/funtimecoding/soil/pkg/kubernetes/argocd/response"
	"github.com/funtimecoding/soil/pkg/kubernetes/constant"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
)

func (c *Client) Applications() ([]*application.Application, error) {
	var p response.Applications

	if e := c.requester.Notation(
		request.Get(constant.ArgocdApplicationsPath),
		&p,
	); e != nil {
		return nil, e
	}

	return application.NewSlice(p.Items), nil
}
