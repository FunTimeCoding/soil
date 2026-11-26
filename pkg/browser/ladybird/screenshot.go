package ladybird

import (
	"encoding/base64"
	"fmt"
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/browser/types/request"
	"github.com/funtimecoding/soil/pkg/browser/types/response"
)

func (c *Client) Screenshot(context string) ([]byte, error) {
	v, e := decodeResult[response.Capture](
		c.send(constant.CaptureMethod, request.NewTarget(context)),
	)

	if e != nil {
		return nil, e
	}

	b, f := base64.StdEncoding.DecodeString(v.Image)

	if f != nil {
		return nil, fmt.Errorf("decode screenshot: %w", f)
	}

	return b, nil
}
