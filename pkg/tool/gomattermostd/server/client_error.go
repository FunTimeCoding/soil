package server

import "github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/server"

func clientError(message string) *server.Error {
	return &server.Error{Error: message}
}
