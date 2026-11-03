package failing_writer

import "net/http"

type Writer struct {
	header  http.Header
	failure error
	Writes  int
}
