package connection_tester

import (
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"net/url"
)

func Refused() error {
	return &url.Error{
		Op:  "Get",
		URL: "http://alfa.example:8080/api/items?api_key=secret",
		Err: outage_tester.RefusedDial(),
	}
}
