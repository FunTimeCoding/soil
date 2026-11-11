package module

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/module_symbol"
	"go/types"
)

func moduleBreakages(
	targets map[string]*types.Package,
	symbols []*module_symbol.Symbol,
	modulePath string,
	newModulePath string,
) []*concern.Concern {
	var result []*concern.Concern
	aliases := map[*types.Package]map[string]string{}

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

		if _, known := aliases[target]; !known {
			aliases[target] = publicAliases(target, modulePath, newModulePath)
		}

		if moduleTypeName(
			nameless(symbol.Object.Type()),
			modulePath,
			newModulePath,
		) == publicTypeName(
			moduleTypeName(nameless(found.Type()), modulePath, newModulePath),
			aliases[target],
		) {
			continue
		}

		was := moduleTypeName(symbol.Object.Type(), modulePath, newModulePath)
		now := publicTypeName(
			moduleTypeName(found.Type(), modulePath, newModulePath),
			aliases[target],
		)
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
