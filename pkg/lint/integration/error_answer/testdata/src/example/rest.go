package example

import (
	"example/generated/server"
	"example/pkg/errors/not_found"
	"example/pkg/errors/unavailable"
	"example/pkg/errors/unreachable"
)

type rest struct {
	connected bool
}

func (r *rest) captureFail(e error) *server.ErrorResponse {
	return &server.ErrorResponse{Error: e.Error()}
}

func (r *rest) item(e error) (any, error) {
	if not_found.Is(e) {
		return server.GetItem400JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.GetItem503JSONResponse{Error: e.Error()}, nil
	}

	if !r.connected {
		return server.GetItem503JSONResponse{Error: "item source down"}, nil
	}

	return server.GetItem404JSONResponse{Error: e.Error()}, nil
}

func (r *rest) answerWith(
	e error,
	_ string,
) *server.ErrorResponse {
	return &server.ErrorResponse{Error: e.Error()}
}

func (r *rest) another(e error) (any, error) {
	if unreachable.Is(e) {
		return server.GetOther503JSONResponse(*r.answerWith(e, "")), nil
	}

	return nil, e
}

func (r *rest) other(e error) (any, error) {
	if unreachable.Is(e) {
		return server.GetOther503JSONResponse(*r.captureFail(e)), nil
	}

	return nil, e
}
