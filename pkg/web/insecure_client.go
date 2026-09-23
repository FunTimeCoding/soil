package web

import (
	"crypto/tls"
	"net/http"
	"sync"
)

var insecureTransport = sync.OnceValue(
	func() *http.Transport {
		return &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	},
)

func InsecureClient() *http.Client {
	return &http.Client{Transport: insecureTransport()}
}
