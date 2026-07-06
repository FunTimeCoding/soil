package basic

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/web/detail_error"
)

func parseDetail(
	body []byte,
	status string,
) error {
	var e response.Error

	if json.Unmarshal(body, &e) == nil && e.Message != "" {
		return detail_error.New(e.Message, status)
	}

	var v2 response.V2Error

	if json.Unmarshal(body, &v2) == nil && len(v2.Errors) > 0 {
		title := v2.Errors[0].Title

		if title == "" {
			title = v2.Errors[0].Detail
		}

		if title != "" {
			return detail_error.New(title, status)
		}
	}

	return fmt.Errorf("%s", status)
}
