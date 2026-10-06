package requester

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
	"net/http"
)

func (r *Requester) build(q *request.Request) (*http.Request, error) {
	target := q.Locator

	if target == "" {
		target = r.base.Copy().Path(q.Path).String()
	}

	result, e := http.NewRequest(q.Method, target, bytes.NewReader(q.Body))

	if e != nil {
		return nil, e
	}

	v := result.URL.Query()

	for k, values := range q.Parameters {
		for _, value := range values {
			v.Add(k, value)
		}
	}

	result.URL.RawQuery = v.Encode()

	if q.ContentType != "" {
		result.Header.Set(constant.ContentType, q.ContentType)
	}

	if r.userAgent != "" {
		result.Header.Set(constant.UserAgent, r.userAgent)
	}

	for k, v := range r.header {
		result.Header.Set(k, v)
	}

	for k, v := range q.Header {
		result.Header.Set(k, v)
	}

	if r.authorizer != nil && result.URL.Host == r.ownHost() {
		if e = r.authorizer.Authorize(result); e != nil {
			return nil, e
		}
	}

	return result, nil
}
