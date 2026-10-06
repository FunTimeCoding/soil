package basic

import (
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/requester"
)

func newRequester(base *locator.Locator) *requester.Requester {
	return requester.New(base).
		WithClient(web.LongStallClient()).
		WithRefusal(refusal)
}
