package subscription

import (
	goclaudedConstant "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/watched_request"
)

func New(
	callsign string,
	k *watched_request.Key,
	reference string,
	alias string,
	d goclaudedConstant.Delivery,
) *Subscription {
	return &Subscription{
		Callsign:  callsign,
		Instance:  k.Instance,
		Project:   k.Project,
		Request:   k.Request,
		Reference: reference,
		Alias:     alias,
		Delivery:  string(d),
	}
}
