package server

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"strings"
)

var bashNoise = map[string]bool{
	"":         true,
	"#":        true,
	"for":      true,
	"if":       true,
	"done":     true,
	"fi":       true,
	"else":     true,
	"else:":    true,
	"elif":     true,
	"in":       true,
	"then":     true,
	"do":       true,
	"import":   true,
	"from":     true,
	"with":     true,
	"continue": true,
	"break":    true,
	"return":   true,
	"pass":     true,
	"try:":     true,
	"except:":  true,
	"class":    true,
	"def":      true,
	"func":     true,
	"package":  true,
	"var":      true,
	"const":    true,
	"type":     true,
	"EOF":      true,
	"PYEOF":    true,
	"GOEOF":    true,
	"local":    true,
	"export":   true,
}

func bashGroupKey(command string) string {
	command = strings.TrimSpace(command)

	if command == "" {
		return ""
	}

	parts := strings.Fields(command)

	if len(parts) == 0 {
		return ""
	}

	first := parts[0]

	if bashNoise[first] {
		return ""
	}

	if strings.HasPrefix(first, "//") ||
		strings.HasPrefix(first, "{") ||
		strings.HasPrefix(first, "(") ||
		strings.HasPrefix(first, ")") ||
		strings.HasPrefix(first, "}") ||
		strings.HasPrefix(first, "\"") ||
		strings.HasPrefix(first, "'") ||
		strings.HasPrefix(first, "s/") ||
		strings.Contains(first, "=") {
		return ""
	}

	if constant.MultiWordPrefixes[first] && len(parts) > 1 {
		return fmt.Sprintf("%s %s", first, parts[1])
	}

	return first
}
