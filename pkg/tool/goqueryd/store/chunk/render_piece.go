package chunk

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func renderPiece(
	t *table,
	header string,
	marker string,
	rows string,
) string {
	var context []string

	for _, s := range []string{t.heading, t.caption, marker} {
		if s != "" {
			context = append(context, s)
		}
	}

	return join.Empty(strings.Join(context, "\n"), "\n\n", header, rows)
}
