package client

import (
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/requester"
	"net/url"
	"strconv"
)

func newRequester(issuer string) *requester.Requester {
	u, e := url.Parse(issuer)

	if e != nil {
		return requester.New(locator.New(issuer))
	}

	l := locator.New(u.Hostname())

	if p, f := strconv.Atoi(u.Port()); f == nil {
		l.Port(p)
	}

	if u.Scheme == "http" {
		l.Insecure()
	}

	return requester.New(l)
}
