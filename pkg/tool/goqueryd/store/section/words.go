package section

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"strings"
)

func Words(s *Section) []string {
	var result []string

	for _, b := range s.Blocks {
		if b.Kind == constant.BlockHeading {
			continue
		}

		for _, w := range strings.Fields(b.Text) {
			if t := strings.Trim(w, constant.MarkupCharacters); t != "" {
				result = append(result, t)
			}
		}
	}

	return result
}
