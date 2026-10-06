package request

import "net/http"

func Absolute(locator string) *Request {
	result := New(http.MethodGet, "")
	result.Locator = locator

	return result
}
