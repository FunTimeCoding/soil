package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
)

func moduleRenamed(
	modulePath string,
	newModulePath string,
	symbols int,
	breakages int,
) *concern.Concern {
	return concern.NewFile(
		constant.ConcernRenamed,
		fmt.Sprintf(
			"%s → %s, %d symbols checked, %d broken",
			modulePath,
			newModulePath,
			symbols,
			breakages,
		),
		"",
		true,
	)
}
