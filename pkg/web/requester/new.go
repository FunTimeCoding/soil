package requester

import (
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func New(base *locator.Locator) *Requester {
	return &Requester{
		base:    base,
		client:  web.StallClient(),
		header:  map[string]string{},
		backoff: constant.Backoff,
	}
}
