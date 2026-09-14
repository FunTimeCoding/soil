package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
)

func moduleRenamed(
	modulePath string,
	newModulePath string,
	symbols int,
	breakages int,
) *concern.Concern {
	return concern.NewFile(
		"renamed",
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
