package report

import "github.com/funtimecoding/soil/pkg/tool/gosourced/constant"

func remap(
	text string,
	mapping map[string]string,
) string {
	if len(mapping) == 0 {
		return text
	}

	return constant.QualifiedReference.ReplaceAllStringFunc(
		text,
		func(reference string) string {
			if to, moved := mapping[reference]; moved {
				return to
			}

			return reference
		},
	)
}
