package gofix

import "fmt"

func keeperReason(
	keeper string,
	reason string,
) string {
	if keeper == "" {
		return fmt.Sprintf("none keeps the name: %s", reason)
	}

	return fmt.Sprintf("%s keeps the name: %s", keeper, reason)
}
