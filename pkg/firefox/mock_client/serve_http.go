package mock_client

import "net/http"

func (c *Client) ServeHTTP(
	w http.ResponseWriter,
	_ *http.Request,
) {
	w.WriteHeader(http.StatusNotImplemented)
}
