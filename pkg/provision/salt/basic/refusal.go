package basic

import (
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"net/http"
	"strings"
)

func refusal(
	status int,
	body []byte,
) error {
	if status == http.StatusNotFound {
		return nil
	}

	s := string(body)
	start := strings.Index(s, "<p>")
	end := strings.Index(s, "</p>")

	if start < 0 || end <= start {
		return nil
	}

	return unexpected.Format(
		"salt-api status: %d: %s",
		status,
		s[start+len("<p>"):end],
	)
}
