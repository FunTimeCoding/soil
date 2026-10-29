package reporter

import "io"

func (r *Reporter) WithOutput(w io.Writer) *Reporter {
	r.output = w
	r.color = false

	return r
}
