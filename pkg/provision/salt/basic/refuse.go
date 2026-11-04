package basic

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func refuse(
	status int,
	body []byte,
) error {
	if status == http.StatusNotFound {
		return nil
	}

	found := paragraph(body)
	reason := found

	if reason == "" {
		reason = http.StatusText(status)
	}

	text := fmt.Sprintf("salt-api status: %d: %s", status, reason)

	if e := refusal.Status(status, text); e != nil {
		return e
	}

	if found == "" {
		return nil
	}

	return unexpected.Format("%s", text)
}
