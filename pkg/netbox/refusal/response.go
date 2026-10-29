package refusal

import (
	"fmt"
	"net/http"
)

func Response(
	status int,
	body []byte,
) error {
	return New(
		status,
		reason(body, fmt.Sprintf("%d %s", status, http.StatusText(status))),
	)
}
