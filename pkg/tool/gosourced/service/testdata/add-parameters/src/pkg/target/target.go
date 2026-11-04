package target

type Answerer interface {
	Reply(message string) string
}

type Server struct{}

func (s *Server) Reply(message string) string {
	return message
}

func Plain(message string) string {
	return message
}

func Grouped(first, second int) int {
	return first + second
}

func Unbound(message string) string {
	return message
}

func Mistyped(message string) string {
	return message
}

func Narrowed(message string) string {
	return message
}

func Valued(message string) string {
	return message
}

func Shadow(message string) string {
	x := message

	return x
}

func Log(prefix string, values ...string) string {
	return prefix + values[0]
}

func Unnamed(string) string {
	return ""
}
