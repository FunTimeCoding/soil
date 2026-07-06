package locator

import "github.com/funtimecoding/soil/pkg/integers"

func (l *Locator) SetInteger(
	k string,
	v int,
) *Locator {
	l.value.Set(k, integers.ToString(v))

	return l
}
