package request

import "net/url"

type Request struct {
	Method      string
	Path        string
	Locator     string
	Parameters  url.Values
	Header      map[string]string
	Body        []byte
	ContentType string
}
