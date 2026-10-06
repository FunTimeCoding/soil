package chunk

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
)

func splitTable(
	content string,
	t *table,
	c face.TokenCounter,
) []Chunk {
	header := content[t.start:t.headerEnd]
	total := len(t.rows)
	widest := fmt.Sprintf(
		constant.TablePieceMarker,
		total,
		total,
		total,
		total,
		total,
	)
	rows := func(from, to int) string {
		return content[t.rows[from].start:t.rows[to-1].end]
	}
	var groups [][2]int

	for from := 0; from < total; {
		to := from + 1

		for to < total &&
			c.Count(renderPiece(t, header, widest, rows(from, to+1))) <=
				c.Allowance() {
			to++
		}

		groups = append(groups, [2]int{from, to})
		from = to
	}

	result := make([]Chunk, len(groups))

	for k, g := range groups {
		position := t.rows[g[0]].start

		if k == 0 {
			position = t.start
		}

		result[k] = Chunk{
			Text: renderPiece(
				t,
				header,
				fmt.Sprintf(
					constant.TablePieceMarker,
					k+1,
					len(groups),
					g[0]+1,
					g[1],
					total,
				),
				rows(g[0], g[1]),
			),
			Position: position,
			Length:   t.rows[g[1]-1].end - position,
		}
	}

	return result
}
