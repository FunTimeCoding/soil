package target

var registry = map[string]int{}

func Register(name string) {
	registry[name] = len(registry)
}

func Count() int {
	return len(registry)
}
