package requester

import "net/http"

func succeeded(status int) bool {
	return status >= http.StatusOK && status < http.StatusMultipleChoices
}
