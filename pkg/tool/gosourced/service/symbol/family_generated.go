package symbol

import (
	"fmt"
	"go/token"
	"go/types"
	"path/filepath"
	"slices"
	"strings"
)

func familyGenerated(
	set *token.FileSet,
	members []*types.Func,
) []string {
	var result []string

	for _, f := range members {
		position := set.Position(f.Pos())
		segments := strings.Split(filepath.Dir(position.Filename), "/")

		if !slices.Contains(segments, "generated") {
			continue
		}

		result = append(
			result,
			fmt.Sprintf(
				"%s.%s is generated code - change its source instead",
				types.TypeString(receiverNamed(f), nil),
				f.Name(),
			),
		)
	}

	return result
}
