package keychain

import (
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
)

func Find(
	service string,
	account string,
) (string, bool, error) {
	r := run.New().NoPanic()
	result := r.Start(
		constant.KeychainCommand,
		constant.KeychainFind,
		"-s",
		service,
		"-a",
		account,
		"-w",
	)

	if r.Exit == constant.KeychainNotFoundExit {
		return "", false, nil
	}

	if r.Error != nil {
		return "", false, r.Error
	}

	return result, true, nil
}
