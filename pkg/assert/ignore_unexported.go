package assert

import (
	"github.com/google/go-cmp/cmp"
	"unicode"
	"unicode/utf8"
)

func ignoreUnexported() cmp.Option {
	return cmp.FilterPath(
		func(p cmp.Path) bool {
			f, okay := p.Index(-1).(cmp.StructField)

			if !okay {
				return false
			}

			r, _ := utf8.DecodeRuneInString(f.Name())

			return !unicode.IsUpper(r)
		},
		cmp.Ignore(),
	)
}
