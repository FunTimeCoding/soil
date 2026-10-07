package language

import "github.com/funtimecoding/soil/pkg/measure/types/language_block"

type Language struct {
	Name          string
	Extensions    []string
	Filenames     []string
	Suffixes      []string
	Shebangs      []string
	LineComments  []string
	BlockComments []*language_block.Block
	Quotes        []string
	RawQuotes     []string
	Nested        bool
}
