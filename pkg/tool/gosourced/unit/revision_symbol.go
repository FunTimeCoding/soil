package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/revision"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/symbol"
	"testing"
)

func revisionSymbol(
	t *testing.T,
	r *revision.Revision,
	key string,
) *symbol.Symbol {
	t.Helper()

	for _, s := range r.Symbols {
		if s.Key() == key {
			return s
		}
	}

	t.Fatalf("no symbol %q in %v", key, revisionKeys(r))

	return nil
}
