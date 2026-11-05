package refusal

import (
	errorsConstant "github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func New(
	status int,
	reason string,
) error {
	if e := refusal.NewTable().With(
		http.StatusBadRequest,
		errorsConstant.ValidationClass,
	).Classify(status, reason); e != nil {
		return e
	}

	return unexpected.Format("netbox status %d: %s", status, reason)
}
