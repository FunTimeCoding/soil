package common

import (
	"encoding/json"
	"errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
	"github.com/netbox-community/go-netbox/v4"
)

func ExtractMessage(e error) (string, bool) {
	var f *netbox.GenericOpenAPIError

	if !errors.As(e, &f) {
		return "", false
	}

	var r response.Error

	if json.Unmarshal(f.Body(), &r) == nil && r.Detail != "" {
		return r.Detail, true
	}

	if message := parseValidationErrors(f.Body()); message != "" {
		return message, true
	}

	return "", false
}
