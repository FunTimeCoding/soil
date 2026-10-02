package goflow

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goflow/constant"
)

func bind(units []string) []string {
	var result []string

	for _, unit := range units {
		if len(result) > 0 && constant.BlockOpenerPattern.MatchString(unit) {
			result[len(result)-1] = join.Space(result[len(result)-1], unit)

			continue
		}

		result = append(result, unit)
	}

	return result
}
