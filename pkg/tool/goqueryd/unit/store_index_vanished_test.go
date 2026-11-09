package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/store_tester"
	"os"
	"path/filepath"
	"testing"
)

func TestIndexSkipsListedFileThatIsGone(t *testing.T) {
	s := store_tester.OpenTestStore(t)
	directory := t.TempDir()
	store_tester.WriteFixture(t, directory, "alfa.md", "# Alfa\n\nStays.\n")
	errors.PanicOnError(
		os.Symlink("charlie.md", filepath.Join(directory, "bravo.md")),
	)
	s.AddCollection("test", directory, constant.DefaultGlob)
	result := s.Index("test")
	assert.Integer(t, 1, result.Indexed)
	assert.Integer(t, 1, s.MustStatus().TotalDocuments)
}
