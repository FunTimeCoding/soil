package failing_writer

import "net/http"

func New(failure error) *Writer {
	return &Writer{header: make(http.Header), failure: failure}
}
