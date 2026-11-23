package keychain_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"os"
	"path/filepath"
	"testing"
)

func New(
	t *testing.T,
	exit string,
) *Tester {
	t.Helper()
	d := t.TempDir()
	errors.PanicOnError(
		os.WriteFile(
			filepath.Join(d, constant.KeychainCommand),
			[]byte(
				join.Empty(
					"#!/bin/sh\n",
					"printf '%s\\n' \"$@\" > \"",
					d,
					"/arguments\"\n",
					"cat > \"",
					d,
					"/input\"\n",
					"[ -n \"",
					exit,
					"\" ] && exit ",
					exit,
					"\n",
					"printf 'stored-secret\\n'\n",
				),
			),
			0o700,
		),
	)
	t.Setenv(
		constant.PathEnvironment,
		join.Empty(d, ":", os.Getenv(constant.PathEnvironment)),
	)

	return &Tester{directory: d}
}
