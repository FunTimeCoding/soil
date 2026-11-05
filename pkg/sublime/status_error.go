package sublime

import (
	"fmt"
	errorsConstant "github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func statusError(
	operation string,
	r *http.Response,
) error {
	reason := refusal.Reason(system.ReadAll(r.Body), r.Status)

	if e := refusal.NewTable().With(
		http.StatusBadRequest,
		errorsConstant.ValidationClass,
	).Classify(
		r.StatusCode,
		fmt.Sprintf("%s: %s", operation, reason),
	); e != nil {
		return e
	}

	return unexpected.Format("%s: %d: %s", operation, r.StatusCode, reason)
}
