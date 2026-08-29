package convert

import "github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"

func Identifier(v string) *server.Identifier {
	return &server.Identifier{Identifier: v}
}
