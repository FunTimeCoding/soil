package basic

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/web/detail_error"
	"net/http"
	"strings"
)

func parseDetail(
	body []byte,
	code int,
) error {
	status := fmt.Sprintf("%d %s", code, http.StatusText(code))
	s := string(body)
	start := strings.Index(s, "<p>")
	end := strings.Index(s, "</p>")

	if start >= 0 && end > start {
		return detail_error.New(s[start+3:end], status)
	}

	return fmt.Errorf("%s", status)
}
