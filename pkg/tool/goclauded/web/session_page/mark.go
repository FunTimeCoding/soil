package session_page

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func Mark(identifier string) string {
	return join.Empty(constant.RosterMarkPrefix, identifier)
}
