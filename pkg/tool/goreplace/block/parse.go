package block

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goreplace/constant"
	"slices"
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

		if !slices.Contains(constant.Openings, lines[i]) {
			return nil, fmt.Errorf(constant.StrayText, i+1, lines[i])
		}

		b := &Block{Number: len(result) + 1, Opening: lines[i]}
		next, e := b.head(lines, i+1)

		if e != nil {
			return nil, e
		}

		replace, end, marker := collect(lines, next, constant.ReplaceMarker)

		if marker == "" {
			return nil, fmt.Errorf(constant.MissingReplace, b.Number)
		}

		b.Replace = replace

		if f := b.validate(); f != nil {
			return nil, f
		}

		result = append(result, b)
		i = end
	}

	if len(result) == 0 {
		return nil, fmt.Errorf(constant.NoBlocks)
	}

	return result, nil
}
