package jellyfin

import "net/http"

type Client struct {
	base string
	http *http.Client
}
