package index

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"reflect"
	"strings"
)

func writeShape(
	b *strings.Builder,
	t reflect.Type,
	seen map[reflect.Type]bool,
) {
	switch t.Kind() {
	case reflect.Struct:
		b.WriteString(join.Empty(t.Kind().String(), "{"))

		if seen[t] {
			b.WriteString("}")

			return
		}

		seen[t] = true

		for i := range t.NumField() {
			f := t.Field(i)
			b.WriteString(join.Empty(f.Name, " `", string(f.Tag), "` "))
			writeShape(b, f.Type, seen)
		}

		b.WriteString("}")
	case reflect.Map:
		b.WriteString(join.Empty(t.Kind().String(), "["))
		writeShape(b, t.Key(), seen)
		b.WriteString("]")
		writeShape(b, t.Elem(), seen)
	case reflect.Pointer, reflect.Slice, reflect.Array:
		b.WriteString(join.Empty(t.Kind().String(), " "))
		writeShape(b, t.Elem(), seen)
	default:
		b.WriteString(join.Empty(t.String(), ";"))
	}
}
