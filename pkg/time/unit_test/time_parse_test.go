package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/time"
	"testing"
)

func TestParse(t *testing.T) {
	result := time.Parse(time.DateYear, "2019-01-01")
	assert.Integer(t, 2019, result.Year())
	assert.Integer(t, 1, int(result.Month()))
	assert.Integer(t, 1, result.Day())
}
