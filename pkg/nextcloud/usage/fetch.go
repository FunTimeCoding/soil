package usage

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/nextcloud/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
)

func (c *Client) Fetch() (*Usage, error) {
	r := web.NewGet(join.Empty(c.base, constant.InformationPath))
	r.Header.Set(constant.TokenHeader, c.token)
	response, e := web.Client().Do(r)

	if e != nil {
		return nil, e
	}

	defer errors.PanicClose(response.Body)

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"server information status: %d",
			response.StatusCode,
		)
	}

	p := &payload{}

	if e = json.NewDecoder(response.Body).Decode(p); e != nil {
		return nil, e
	}

	return &Usage{
		Files:  p.Wrapper.Body.Nextcloud.Storage.Files,
		Shares: p.Wrapper.Body.Nextcloud.Shares.Count,
	}, nil
}
