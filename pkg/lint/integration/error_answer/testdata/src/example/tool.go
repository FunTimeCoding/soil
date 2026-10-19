package example

import (
	"errors"
	"example/generative/mark/response"
	"example/pkg/errors/not_found"
	"example/pkg/errors/unreachable"
	"example/pkg/errors/validation"
)

type tool struct{}

func (t *tool) captureFail(e error) (string, error) {
	return e.Error(), nil
}

func (t *tool) answered(e error) (string, error) {
	if not_found.Is(e) {
		return response.Fail("%s", e)
	}

	if validation.Is(e) {
		return t.captureFail(e)
	}

	if unreachable.Is(e) || not_found.Is(e) {
		return response.Fail("%s", e)
	}

	if !unreachable.Is(e) {
		return response.Fail("%s", e)
	}

	switch {
	case not_found.Is(e):
		return response.Fail("%s", e)
	}

	if v, okay := errors.AsType[*validation.Detail](e); okay {
		return t.captureFail(v)
	}

	var d *validation.Detail

	if errors.As(e, &d) {
		return response.Fail("%s", d.Message)
	}

	return t.captureFail(e)
}
