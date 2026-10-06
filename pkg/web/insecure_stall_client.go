package web

import "net/http"

func InsecureStallClient() *http.Client {
	return &http.Client{Transport: insecureStallTransport()}
}
