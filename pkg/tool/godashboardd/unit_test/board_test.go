package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board"
	"testing"
)

func TestLoad(t *testing.T) {
	b := board.Load(fixture.Path(constant.BoardPath, "board.yaml"))
	assert.String(t, "prometheus.example", b.Connection.Prometheus.Host)
	assert.Integer(t, 2, len(b.Top))
	assert.String(t, "monitoring", b.Top[0].Name)
	assert.String(t, "Monitor", b.Top[0].Sections[0].Name)
	assert.Integer(t, 2, len(b.Top[0].Sections[0].Entries))
	assert.String(
		t,
		"count(up == 1)",
		b.Top[0].Sections[0].Entries[0].Rows[0].Query,
	)
	assert.String(t, "nextcloud", b.Top[1].Sections[0].Entries[0].Widget)
	assert.Integer(t, 2, b.Tail.Columns)
	assert.Integer(t, 2, len(b.Tail.Sections))
	assert.String(t, "", b.Tail.Sections[1].Name)
	assert.Integer(t, 6, len(b.Entries()))
}
