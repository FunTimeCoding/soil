package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"go/types"
)

func moduleBreakages(
	targets map[string]*types.Package,
	symbols []*ModuleSymbol,
	modulePath string,
	newModulePath string,
) []*concern.Concern {
	var result []*concern.Concern

	for _, symbol := range symbols {
		targetPath := moduleTargetPath(
			symbol.PackagePath,
			modulePath,
			newModulePath,
		)
		target := targets[targetPath]

		if target == nil {
			result = append(
				result,
				moduleBreakage(
					symbol,
					constant.ConcernMissing,
					fmt.Sprintf("package not found: %s", targetPath),
				),
			)

			continue
		}

		found := moduleMember(target, symbol)

		if found == nil {
			result = append(
				result,
				moduleBreakage(
					symbol,
					constant.ConcernMissing,
					fmt.Sprintf(
						"%s no longer exists in %s",
						moduleMemberName(symbol),
						targetPath,
					),
				),
			)

			continue
		}

		was := moduleTypeName(symbol.Object.Type(), modulePath, newModulePath)
		now := moduleTypeName(found.Type(), modulePath, newModulePath)

		if was == now {
			continue
		}

		result = append(
			result,
			moduleBreakage(
				symbol,
				constant.ConcernChanged,
				fmt.Sprintf("%s: %s → %s", moduleMemberName(symbol), was, now),
			),
		)
	}

	return result
}
