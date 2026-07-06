package basic

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
	"github.com/funtimecoding/soil/pkg/web/detail_error"
)

func parseDetail(
	body []byte,
	status string,
) error {
	var e response.Error

	if json.Unmarshal(body, &e) == nil && e.Detail != "" {
		return detail_error.New(e.Detail, status)
	}

	return fmt.Errorf("%s", status)
}
