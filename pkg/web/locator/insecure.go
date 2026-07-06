package locator

import "github.com/funtimecoding/soil/pkg/web/constant"

func (l *Locator) Insecure() *Locator {
	l.scheme = constant.Insecure

	return l
}
