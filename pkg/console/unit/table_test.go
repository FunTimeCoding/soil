package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/console/table"
	"testing"
)

func TestTableWidthsAndAlignment(t *testing.T) {
	v := table.New("NAME", "COUNT").Right(1)
	v.Add("alfa", "1")
	v.Add("bravo-charlie", "1234")
	assert.String(
		t,
		"NAME           COUNT\nalfa               1\nbravo-charlie   1234\n",
		v.Render(),
	)
}

func TestTableShortRow(t *testing.T) {
	v := table.New("A", "B", "C")
	v.Add("x")
	assert.String(t, "A  B  C\nx      \n", v.Render())
}

func TestTableHeaderOnly(t *testing.T) {
	assert.String(t, "ONLY\n", table.New("ONLY").Render())
}
