package band

import (
	"io"
	"net/http"
)

func read(response *http.Response) ([]byte, error) {
	body, e := io.ReadAll(response.Body)

	if f := response.Body.Close(); f != nil {
		return nil, f
	}

	return body, e
}
