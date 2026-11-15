package upstream_server

import "net/http"

func bodyAllowed(status int) bool {
	return status >= http.StatusOK &&
		status != http.StatusNoContent &&
		status != http.StatusNotModified
}
