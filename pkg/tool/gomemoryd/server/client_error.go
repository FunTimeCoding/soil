package server

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/server"

func clientError(e error) *server.Error {
	return &server.Error{Error: e.Error()}
}
