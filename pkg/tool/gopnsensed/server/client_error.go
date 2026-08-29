package server

import "github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"

func clientError(message string) *server.Error {
	return &server.Error{Error: message}
}
