package band

import (
	"io"
	"net/http"
)

func drain(response *http.Response) error {
	_, e := io.Copy(io.Discard, response.Body)

	if f := response.Body.Close(); f != nil {
		return f
	}

	return e
}
