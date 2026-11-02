package constant

import "github.com/funtimecoding/soil/pkg/constant"

const (
	NoLimit int = -1

	DictionaryFile = "dictionary.dic"
	MinimumStem    = 3

	WikiPathEnvironment = "WIKI_PATH"
)

const (
	DiffEqual  = " "
	DiffDelete = "-"
	DiffInsert = "+"
	DiffGap    = "@@"
)

var (
	DictionarySkip   = map[string]bool{".git": true, "tmp": true}
	DictionaryPrefix = map[string]bool{
		"Containerfile.": true,
		"Dockerfile.":    true,
	}
	DictionaryExtension = map[string]bool{
		".conf":      true,
		".envrc":     true,
		".gitignore": true,
		".go":        true,
		".json":      true,
		".md":        true,
		".rego":      true,
		".sh":        true,
		".txt":       true,
		".xml":       true,
		".yaml":      true,
		".yml":       true,
	}
	DictionaryNoExtension = map[string]bool{
		constant.ContainerFile: true,
		constant.DockerFile:    true,
	}
)

var (
	Inflections = map[string][]string{
		"ies":  {"y"},
		"es":   {""},
		"s":    {""},
		"ed":   {"", "e"},
		"ied":  {"y"},
		"ing":  {"", "e"},
		"er":   {"", "e"},
		"ier":  {"y"},
		"est":  {"", "e"},
		"iest": {"y"},
	}
	Doubling = map[string]bool{"ed": true, "ing": true, "er": true, "est": true}
)
