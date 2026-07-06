package client_context

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"strings"
)

func Current() string {
	return strings.TrimSpace(
		system.Run(
			constant.Kubectl,
			constant.Configuration,
			constant.CurrentContext,
		),
	)
}
