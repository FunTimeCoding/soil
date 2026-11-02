package band

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/band/response"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"strings"
)

func faulted(
	resource string,
	f response.Fault,
) error {
	text := fmt.Sprintf("band %s: fault %s", resource, f.Text)
	_, local, _ := strings.Cut(f.Subcode, ":")

	switch local {
	case constant.FaultAccessDenied:
		return forbidden.Format("%s", text)
	case constant.FaultAlreadyExists, constant.FaultConcurrency:
		return conflict.Format("%s", text)
	case constant.FaultQuotaLimit,
		constant.FaultEndpointUnavailable,
		constant.FaultTimedOut:

		return unavailable.Format("%s", text)
	}

	return unexpected.Format("%s", text)
}
