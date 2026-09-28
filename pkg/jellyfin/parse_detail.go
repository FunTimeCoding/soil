package jellyfin

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/jellyfin/response"
	"github.com/funtimecoding/soil/pkg/web/detail_error"
)

func parseDetail(
	body []byte,
	status string,
) error {
	var e response.Error

	if json.Unmarshal(body, &e) == nil && e.Title != "" {
		return detail_error.New(e.Title, status)
	}

	if len(body) > 0 {
		return detail_error.New(string(body), status)
	}

	return fmt.Errorf("%s", status)
}
