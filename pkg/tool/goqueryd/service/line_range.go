package service

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/chunk"
	"strings"
)

func lineRange(
	body string,
	c chunk.Chunk,
) (int, int) {
	first := strings.Count(body[:c.Position], stringsConstant.Unix) + 1

	return first, first + strings.Count(
		strings.TrimRight(body[c.Position:c.Position+c.Length], stringsConstant.Unix),
		stringsConstant.Unix,
	)
}
