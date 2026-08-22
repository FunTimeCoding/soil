package basic

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/technitium/envelope"
	"github.com/funtimecoding/soil/pkg/web/detail_error"
)

func parseDetail(
	body []byte,
	status string,
) error {
	var e envelope.Envelope

	if json.Unmarshal(body, &e) == nil && e.Message != "" {
		return detail_error.New(e.Message, status).WithBody(body)
	}

	if len(body) > 0 {
		return detail_error.New(string(body), status).WithBody(body)
	}

	return fmt.Errorf("%s", status)
}
