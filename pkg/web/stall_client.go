package web

import (
	"crypto/tls"
	"net/http"
	"sync"
)

var stallTransport = sync.OnceValue(
	func() *http.Transport {
		return newStallTransport(nil)
	},
)

var insecureStallTransport = sync.OnceValue(
	func() *http.Transport {
		return newStallTransport(&tls.Config{InsecureSkipVerify: true})
	},
)

func StallClient() *http.Client {
	return &http.Client{Transport: stallTransport()}
}
