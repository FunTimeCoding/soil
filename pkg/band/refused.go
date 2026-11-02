package band

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func refused(
	resource string,
	code int,
	reason string,
) error {
	text := fmt.Sprintf("band %s: %s", resource, reason)

	if e := refusal.Status(code, text); e != nil {
		return e
	}

	return unexpected.Format("%s", text)
}
