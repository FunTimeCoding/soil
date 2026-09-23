package channel

type Sink interface {
	Push(content string, meta map[string]string)
}
