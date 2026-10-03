package service_tester

import (
	"github.com/funtimecoding/soil/pkg/git"
	"path/filepath"
)

func ServiceTestdata(name string) string {
	return filepath.Join(
		git.FindDirectory(),
		"pkg/tool/gosourced/service/testdata",
		name,
	)
}
