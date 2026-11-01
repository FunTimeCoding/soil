package client_page

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
)

func Path(identifier string) string {
	return join.Empty(constant.ClientsPrefix, identifier)
}
