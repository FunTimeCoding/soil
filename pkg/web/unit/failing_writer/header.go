package failing_writer

import "net/http"

func (w *Writer) Header() http.Header {
	return w.header
}
