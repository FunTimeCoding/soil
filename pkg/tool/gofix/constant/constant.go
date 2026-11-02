package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"gofix",
	"Go source fixer",
	"gofix [flags] [pattern...]",
)

const ConcernConcurrentWrite = "concurrent_write"
const (
	MaxSingleParameterLength = 80
	MaxFormatPasses          = 10
	DiffContext              = 2
)

var (
	BuiltinNames = map[string]bool{
		"bool":       true,
		"byte":       true,
		"complex64":  true,
		"complex128": true,
		"error":      true,
		"float32":    true,
		"float64":    true,
		"int":        true,
		"int8":       true,
		"int16":      true,
		"int32":      true,
		"int64":      true,
		"rune":       true,
		"string":     true,
		"uint":       true,
		"uint8":      true,
		"uint16":     true,
		"uint32":     true,
		"uint64":     true,
		"uintptr":    true,
		"any":        true,
		"comparable": true,
		"append":     true,
		"cap":        true,
		"clear":      true,
		"close":      true,
		"complex":    true,
		"copy":       true,
		"delete":     true,
		"imag":       true,
		"len":        true,
		"make":       true,
		"max":        true,
		"min":        true,
		"new":        true,
		"panic":      true,
		"print":      true,
		"println":    true,
		"real":       true,
		"recover":    true,
		"true":       true,
		"false":      true,
		"iota":       true,
		"nil":        true,
	}
)
