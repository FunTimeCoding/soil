package argocd

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/argocd/application"
	"github.com/funtimecoding/soil/pkg/argocd/constant"
	"github.com/funtimecoding/soil/pkg/argocd/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
)

func (c *Client) Applications() ([]*application.Application, error) {
	r := web.NewGet(join.Empty(c.base, constant.ApplicationsPath))
	web.Bearer(r, c.token)
	reply, e := web.Client().Do(r)

	if e != nil {
		return nil, e
	}

	defer errors.PanicClose(reply.Body)

	if reply.StatusCode != http.StatusOK {
		return nil, unexpected.Format(
			"applications status: %d",
			reply.StatusCode,
		)
	}

	var p response.Applications

	if e = json.NewDecoder(reply.Body).Decode(&p); e != nil {
		return nil, e
	}

	return application.NewSlice(p.Items), nil
}
