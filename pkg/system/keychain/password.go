package keychain

import (
	"github.com/funtimecoding/soil/pkg/system/run"
	"strings"
)

func Password(service string) string {
	return strings.TrimSpace(
		run.New().Start(
			"security",
			"find-generic-password",
			"-s",
			service,
			"-w",
		),
	)
}
