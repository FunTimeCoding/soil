package callback_report

import (
	provisionConstant "github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/types/host_result"
)

func hostResult(
	host string,
	o *Outcome,
) *host_result.Result {
	if o.Unreachable {
		return host_result.New(
			host,
			provisionConstant.HostUnreachable,
			o.ReturnCode,
			o.Stdout,
			o.Message,
		)
	}

	status := provisionConstant.StoreStatusSuccess

	if o.Failed {
		status = provisionConstant.StoreStatusError
	}

	return host_result.New(host, status, o.ReturnCode, o.Stdout, o.Stderr)
}
