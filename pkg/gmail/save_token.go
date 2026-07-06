package gmail

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors"
	"golang.org/x/oauth2"
	"os"
)

func (c *Client) saveToken(
	t *oauth2.Token,
	path string,
) {
	f, e := os.OpenFile(
		path,
		os.O_RDWR|os.O_CREATE|os.O_TRUNC,
		0600,
	)
	errors.PanicOnError(e)
	defer errors.PanicClose(f)
	errors.PanicOnError(json.NewEncoder(f).Encode(t))
}
