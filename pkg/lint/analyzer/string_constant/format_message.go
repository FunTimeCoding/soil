package string_constant

import (
	"fmt"
	"strings"
)

func formatMessage(
	value string,
	list []knownConstant,
) string {
	if len(list) == 1 {
		return fmt.Sprintf(
			"string literal %q has constant %s.%s",
			value,
			list[0].packageName,
			list[0].name,
		)
	}

	names := make([]string, len(list))

	for i, c := range list {
		names[i] = fmt.Sprintf("%s.%s", c.packageName, c.name)
	}

	return fmt.Sprintf(
		"string literal %q has constants %s",
		value,
		strings.Join(names, " or "),
	)
}
