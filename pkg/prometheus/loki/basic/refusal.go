package basic

import (
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"net/http"
	"strings"
)

func refusal(
	status int,
	body []byte,
) error {
	text := strings.TrimSpace(string(body))

	if status == http.StatusNotFound || text == "" ||
		strings.ContainsAny(text[:1], constant.LokiMarkupOpenings) {
		return nil
	}

	return unexpected.Format("loki status: %d: %s", status, text)
}
