package locator

import (
	"github.com/funtimecoding/soil/pkg/errors"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"net/url"
	"strings"
)

func (l *Locator) Target(t string) *Locator {
	path, query, _ := strings.Cut(t, stringsConstant.Question)
	l.Path(path)
	v, e := url.ParseQuery(query)
	errors.PanicOnError(e)

	for k, values := range v {
		for _, value := range values {
			l.Add(k, value)
		}
	}

	return l
}
