package revision_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/revision"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/symbol"
)

func (o *Tester) Symbol(
	r *revision.Revision,
	key string,
) *symbol.Symbol {
	o.t.Helper()

	for _, s := range r.Symbols {
		if s.Key() == key {
			return s
		}
	}

	o.t.Fatalf("no symbol %q in %v", key, Keys(r))

	return nil
}
