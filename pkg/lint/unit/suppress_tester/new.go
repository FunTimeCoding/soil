package suppress_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"go/parser"
	"go/token"
	"testing"
)

func New(
	t *testing.T,
	source string,
) *Tester {
	t.Helper()
	s := token.NewFileSet()
	f, e := parser.ParseFile(s, "example.go", source, parser.ParseComments)
	assert.FatalOnError(t, e)

	return &Tester{t: t, fileSet: s, file: f}
}
