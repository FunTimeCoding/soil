package aptly

import (
	"bytes"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func (c *Client) Upload(
	directory string,
	filePath string,
) error {
	f, e := os.Open(filePath)

	if e != nil {
		return e
	}

	defer errors.LogClose(f)
	b := &bytes.Buffer{}
	writer := multipart.NewWriter(b)
	part, g := writer.CreateFormFile("file", filepath.Base(filePath))

	if g != nil {
		return g
	}

	if _, h := io.Copy(part, f); h != nil {
		return h
	}

	if i := writer.Close(); i != nil {
		return i
	}

	_, j := c.requester.Bytes(
		request.New(
			http.MethodPost,
			fmt.Sprintf("/api/files/%s", directory),
		).WithBody(writer.FormDataContentType(), b.Bytes()),
	)

	return j
}
