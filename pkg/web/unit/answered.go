package unit

import "net/http"

func answered(status int) *http.Response {
	return &http.Response{StatusCode: status, Status: http.StatusText(status)}
}
