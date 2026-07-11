package target

import (
	"github.com/funtimecoding/soil/pkg/source/example/probe"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func Report() string {
	v := probe.Probe{Label: "probe"}

	return join.Empty(v.Tag(), v.Describe())
}
