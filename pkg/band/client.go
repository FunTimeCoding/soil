package band

import "net/http"

type Client struct {
	base     string
	user     string
	password string
	client   *http.Client
}
