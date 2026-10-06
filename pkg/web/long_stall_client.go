package web

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"sync"
)

var longStallTransport = sync.OnceValue(
	func() *http.Transport {
		result := newStallTransport(nil)
		result.ResponseHeaderTimeout = constant.LongResponseHeaderTimeout

		return result
	},
)

func LongStallClient() *http.Client {
	return &http.Client{Transport: longStallTransport()}
}
