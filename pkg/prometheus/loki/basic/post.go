package basic

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"log"
)

func (c *Client) Post(
	l string,
	body []byte,
) {
	if c.verbose {
		console.Format("POST %s\n%s\n", l, body)
	}

	r := web.NewPostBytes(l, bytes.NewReader(body))
	r.Header.Set(constant.ContentType, constant.Object)
	r.SetBasicAuth(c.user, c.password)
	s := web.Send(web.Client(), r)

	if s.StatusCode >= 400 {
		log.Panicf("push failed %d: %s", s.StatusCode, web.ReadString(s))
	}
}
