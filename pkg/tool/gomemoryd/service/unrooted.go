package service

import (
	"github.com/funtimecoding/soil/pkg/errors/not_configured"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
)

func unrooted() error {
	return not_configured.Format(
		"%s is not set - no reference root, nothing checked",
		constant.ReferenceRootEnvironment,
	)
}
