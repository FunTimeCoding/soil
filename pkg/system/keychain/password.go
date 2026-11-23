package keychain

import (
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
	"strings"
)

func Password(service string) string {
	return strings.TrimSpace(
		run.New().Start(
			constant.KeychainCommand,
			constant.KeychainFind,
			"-s",
			service,
			"-w",
		),
	)
}
