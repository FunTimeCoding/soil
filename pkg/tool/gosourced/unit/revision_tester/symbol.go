package revision_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/revision"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/symbol"
	"testing"
)

func Symbol(
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

	t.Fatalf("no symbol %q in %v", key, Keys(r))

	return nil
}
