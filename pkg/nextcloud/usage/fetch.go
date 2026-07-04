package usage

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (c *Client) Fetch() (*Usage, error) {
	r := web.NewGet(join.Empty(c.base, informationPath))
	r.Header.Set(tokenHeader, c.token)
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
