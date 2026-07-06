package locator

import "github.com/funtimecoding/soil/pkg/integers64"

func (l *Locator) SetInteger64(
	k string,
	v int64,
) *Locator {
	l.value.Set(k, integers64.ToString(v))

	return l
}
