package aptly

import "net/http"

type Client struct {
	Base     string
	Username string
	Password string
	client   *http.Client
}
