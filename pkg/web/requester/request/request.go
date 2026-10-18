package request

import (
	"context"
	"net/url"
)

type Request struct {
	Context     context.Context
	Method      string
	Path        string
	Locator     string
	Parameters  url.Values
	Header      map[string]string
	Body        []byte
	ContentType string
}
