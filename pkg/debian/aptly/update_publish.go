package aptly

import (
	"bytes"
	"fmt"
	"github.com/funtimecoding/soil/pkg/debian/aptly/request"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func (c *Client) UpdatePublish(
	distribution string,
	passphraseFile string,
) error {
	r := request.UpdatePublish{Signing: request.NewSignOption(passphraseFile)}
	s, e := c.send(
		constant.Put,
		fmt.Sprintf("/api/publish/:./%s", distribution),
		bytes.NewReader(notation.Marshal(r)),
	)
	errors.PanicOnError(e)
	defer errors.LogClose(s.Body)

	if s.StatusCode != http.StatusOK {
		return unexpected.Format(
			"update publish status: %s %s",
			s.Status,
			string(system.ReadAll(s.Body)),
		)
	}

	return nil
}
