package basic

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func (c *Client) Propfind() {
	r := web.NewPropfind(c.fileRoot)

	if false {
		// WebDAV is XML
		r.Header.Set(constant.Accept, constant.Object)
	}

	r.SetBasicAuth(c.user, c.password)
	s := web.Send(web.Client(), r)
	defer errors.LogClose(s.Body)

	switch s.StatusCode {
	case http.StatusMultiStatus:
		console.Line("success")

		if false {
			// A lot of XML
			console.Line("response body:", web.ReadString(s))
		}
	case http.StatusUnauthorized:
		console.Line(constant.Unauthorized)
	default:
		console.Format("unexpected status: %d\n", s.StatusCode)
	}
}
