package aptly

import (
	"bytes"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"mime/multipart"
	"net/http"
	"path/filepath"
)

func (c *Client) Upload(
	directory string,
	filePath string,
) error {
	f := system.Open(filePath)
	defer errors.LogClose(f)
	b := &bytes.Buffer{}
	writer := multipart.NewWriter(b)
	part, e := writer.CreateFormFile("file", filepath.Base(filePath))
	errors.PanicOnError(e)
	system.Copy(f, part)
	errors.PanicClose(writer)
	r, g := http.NewRequest(
		constant.Post,
		fmt.Sprintf("%s/api/files/%s", c.Base, directory),
		b,
	)
	errors.PanicOnError(g)

	if c.Username != "" {
		r.SetBasicAuth(c.Username, c.Password)
	}

	r.Header.Set(constant.ContentType, writer.FormDataContentType())
	s, i := c.client.Do(r)
	errors.PanicOnError(i)
	defer errors.LogClose(s.Body)

	if s.StatusCode != http.StatusOK {
		return unexpected.Format(
			"upload status: %s %s",
			s.Status,
			string(system.ReadAll(s.Body)),
		)
	}

	return nil
}
