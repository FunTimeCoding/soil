package failing_writer

func (w *Writer) Write([]byte) (int, error) {
	w.Writes++

	return 0, w.failure
}
