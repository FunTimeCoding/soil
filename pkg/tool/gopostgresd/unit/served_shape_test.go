package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/postgres/types/result"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"testing"
)

func TestTheModelContextAnswersColumnsAndRows(t *testing.T) {
	s, _, _, m := served(t, nil)
	m.SeedResult(
		"SELECT a.id, b.id",
		result.New([]string{"id", "id"}, [][]any{{1, 7}}),
	)
	c := s.Connect(t)
	defer c.Close()
	var v wireResult
	notation.MustDecodeBytes(
		[]byte(
			c.MustCallTool(
				constant.Query,
				map[string]any{"sql": "SELECT a.id, b.id"},
			),
		),
		&v,
		false,
	)
	assert.Count(t, 2, v.Columns)
	assert.Any(t, 7.0, v.Rows[0][1])
}
