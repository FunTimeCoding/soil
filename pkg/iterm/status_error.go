package iterm

import (
	"fmt"
	errorsConstant "github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func statusError(
	operation string,
	r *http.Response,
) error {
	text := reason(r)

	if e := refusal.NewTable().With(
		http.StatusUnprocessableEntity,
		errorsConstant.ValidationClass,
	).Classify(
		r.StatusCode,
		fmt.Sprintf("%s: %s", operation, text),
	); e != nil {
		return e
	}

	return unexpected.Format("%s: %d: %s", operation, r.StatusCode, text)
}
