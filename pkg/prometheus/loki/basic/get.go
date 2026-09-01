package basic

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) Get(l string) string {
	if c.verbose {
		console.Format("GET %s\n", l)
	}

	r := web.NewGet(l)
	r.SetBasicAuth(c.user, c.password)
	s := web.Send(web.Client(), r)

	if c.verbose {
		console.Line("Request:")
		console.Line(r)
		console.Line("Response:")
		console.Line(s)

		if s.StatusCode >= 400 {
			console.Format("Error: %s\n", web.ReadString(s))

			return ""
		}
	}

	return web.ReadString(s)
}
