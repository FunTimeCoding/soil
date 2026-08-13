package request_context

import "net/http"

func (c *Context) Redirect(target string) {
	if target == "" {
		target = c.Referer()
	}

	http.Redirect(c.writer, c.request, target, http.StatusFound)
}
