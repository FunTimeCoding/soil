package basic

import (
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/requester"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
)

func newRequester(base *locator.Locator) *requester.Requester {
	return requester.New(base).
		WithClient(stall_client.New().Long().Build().Client()).
		WithRefusal(refusal)
}
