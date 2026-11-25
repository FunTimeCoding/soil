package keychain

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
	"strings"
)

func Save(
	service string,
	account string,
	secret string,
) error {
	if strings.ContainsAny(
		join.Empty(service, account, secret),
		constant.KeychainUnsafe,
	) {
		return validation.New(
			"keychain entry holds a quote, backslash or newline: %s",
			service,
		)
	}

	command := fmt.Sprintf(
		"add-generic-password -U -s \"%s\" -a \"%s\" -w \"%s\"",
		service,
		account,
		secret,
	)

	if len(command) > constant.KeychainLineLimit {
		return validation.New(
			"keychain entry for %s is %d bytes, security -i reads %d",
			service,
			len(command),
			constant.KeychainLineLimit,
		)
	}

	r := run.New().NoPanic()
	r.Input = strings.NewReader(join.Empty(command, "\n"))
	r.Start(constant.KeychainCommand, "-i")

	return r.Error
}
