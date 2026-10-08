package block

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goreplace/constant"
	"strings"
)

func Parse(input string) ([]*Block, error) {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	var result []*Block
	i := 0

	for i < len(lines) {
		if strings.TrimSpace(lines[i]) == "" {
			i++

			continue
		}

		if lines[i] != constant.SearchMarker {
			return nil, fmt.Errorf(constant.StrayText, i+1, lines[i])
		}

		b := &Block{Number: len(result) + 1}
		search, next, found := collect(lines, i+1, constant.DividerMarker)

		if !found {
			return nil, fmt.Errorf(constant.MissingDivider, b.Number)
		}

		replace, end, found := collect(lines, next, constant.ReplaceMarker)

		if !found {
			return nil, fmt.Errorf(constant.MissingReplace, b.Number)
		}

		if search == "" {
			return nil, fmt.Errorf(constant.EmptySearch, b.Number)
		}

		b.Search = search
		b.Replace = replace
		result = append(result, b)
		i = end
	}

	if len(result) == 0 {
		return nil, fmt.Errorf(constant.NoBlocks)
	}

	return result, nil
}
