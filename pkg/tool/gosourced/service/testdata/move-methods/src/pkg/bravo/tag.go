package bravo

import "strconv"

func (w *Worker) Tag(v int) string {
	return strconv.Itoa(v)
}
