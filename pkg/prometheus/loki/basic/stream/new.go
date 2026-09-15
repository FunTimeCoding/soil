package stream

func New(labels map[string]string) *Stream {
	return &Stream{Stream: labels}
}
