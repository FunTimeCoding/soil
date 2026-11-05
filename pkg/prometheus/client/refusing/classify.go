package refusing

import (
	"fmt"
	errorsConstant "github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func classify(
	status int,
	body []byte,
) error {
	if status/100 == 2 {
		return nil
	}

	reason := refusal.Reason(body, http.StatusText(status))

	if e := refusal.NewTable().
		Without(http.StatusNotFound).
		Without(http.StatusServiceUnavailable).
		With(http.StatusBadRequest, errorsConstant.ValidationClass).
		With(http.StatusUnprocessableEntity, errorsConstant.ValidationClass).
		Classify(status, fmt.Sprintf("prometheus: %s", reason)); e != nil {
		return e
	}

	return unexpected.Format("prometheus status: %d: %s", status, reason)
}
