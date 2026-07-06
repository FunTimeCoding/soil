package web

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"github.com/funtimecoding/soil/pkg/web/spy_writer"
	"testing"
)

func TestWriteOkay(t *testing.T) {
	w := spy_writer.New()
	assert.Integer(t, 4, WriteOkay(w, upper.Alfa))
	assert.Any(t, []byte("Alfa"), w.Written)
	assert.Integer(t, 200, w.StatusCode)
}
