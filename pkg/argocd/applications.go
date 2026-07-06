package argocd

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/argocd/application"
	"github.com/funtimecoding/soil/pkg/argocd/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
)

func (c *Client) Applications() ([]*application.Application, error) {
	r := web.NewGet(join.Empty(c.base, applicationsPath))
	web.Bearer(r, c.token)
	reply, e := web.Client().Do(r)

	if e != nil {
		return nil, e
	}

	defer errors.PanicClose(reply.Body)

	if reply.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
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
