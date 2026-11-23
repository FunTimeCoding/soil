package state_tree

import (
	"github.com/funtimecoding/soil/pkg/salt/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func stateName(relative string) string {
	result := strings.TrimSuffix(relative, constant.StateExtension)
	result = strings.TrimSuffix(result, join.Empty("/", constant.InitName))

	return strings.ReplaceAll(result, "/", constant.StateSeparator)
}
