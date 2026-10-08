package web

import "net/http"

func HaltingClient() *http.Client {
	result := StallClient()
	result.CheckRedirect = func(
		*http.Request,
		[]*http.Request,
	) error {
		return http.ErrUseLastResponse
	}

	return result
}
