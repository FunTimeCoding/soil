package revision

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/symbol"
	"path"
	"strconv"
)

func distinguish(symbols []*symbol.Symbol) {
	for _, same := range collisions(symbols) {
		for _, s := range same {
			s.Variant = join.Empty(constant.VariantSeparator, path.Base(s.File))
		}
	}

	for _, same := range collisions(symbols) {
		for i, s := range same {
			s.Variant = join.Empty(
				s.Variant,
				constant.OrdinalSeparator,
				strconv.Itoa(i+1),
			)
		}
	}
}
