package fritz

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/funtimecoding/soil/pkg/digest"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/fritz/response"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"io"
	"net/http"
)

func (c *Client) call(
	path string,
	service string,
	action string,
) ([]byte, error) {
	envelope := fmt.Sprintf(
		`<?xml version="1.0" encoding="utf-8"?><s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/" s:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/"><s:Body><u:%s xmlns:u="%s"/></s:Body></s:Envelope>`,
		action,
		service,
	)
	address := join.Empty(c.base, path)
	first, e := http.NewRequest(
		http.MethodPost,
		address,
		bytes.NewBufferString(envelope),
	)

	if e != nil {
		return nil, e
	}

	soapHeader(first, service, action)
	challenge, f := c.client.Do(first)

	if f != nil {
		return nil, f
	}

	if g := drain(challenge); g != nil {
		return nil, g
	}

	if challenge.StatusCode != http.StatusUnauthorized {
		return nil, unexpected.Format(
			"fritz %s: expected digest challenge, got %s",
			path,
			challenge.Status,
		)
	}

	second, g := http.NewRequest(
		http.MethodPost,
		address,
		bytes.NewBufferString(envelope),
	)

	if g != nil {
		return nil, g
	}

	soapHeader(second, service, action)
	second.Header.Set(
		"Authorization",
		digest.Authorization(
			challenge.Header.Get("WWW-Authenticate"),
			http.MethodPost,
			path,
			c.user,
			c.password,
		),
	)
	r, h := c.client.Do(second)

	if h != nil {
		return nil, h
	}

	body, i := io.ReadAll(r.Body)

	if j := r.Body.Close(); j != nil {
		return nil, j
	}

	if i != nil {
		return nil, i
	}

	if r.StatusCode >= http.StatusBadRequest {
		var a response.Fault

		if xml.Unmarshal(body, &a) == nil && a.Code != "" {
			return nil, unexpected.Format(
				"fritz %s %s: fault %s (%s)",
				path,
				action,
				a.Code,
				a.Description,
			)
		}

		return nil, unexpected.Format("fritz %s %s: %s", path, action, r.Status)
	}

	return body, nil
}
