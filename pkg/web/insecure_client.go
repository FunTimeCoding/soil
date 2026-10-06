package web

import "net/http"

func InsecureClient() *http.Client {
	return InsecureStallClient()
}
