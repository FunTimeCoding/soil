package requester

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
)

func New(base *locator.Locator) *Requester {
	return &Requester{
		base:    base,
		client:  stall_client.New().Build().Client(),
		header:  map[string]string{},
		table:   refusal.NewTable(),
		reason:  refusal.Reason,
		backoff: constant.Backoff,
	}
}
