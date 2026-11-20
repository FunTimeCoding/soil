package request

import (
	"context"
	"net/url"
	"time"
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
	Refusal     func(
		int,
		[]byte,
	) error
	Statuses map[int]string
	Since    time.Time
}
