package resolve

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func TagFlags(tags []string) []string {
	if len(tags) == 0 {
		return nil
	}

	return []string{fmt.Sprintf("-tags=%s", join.Comma(tags))}
}
