package language

type Language struct {
	Name          string
	Extensions    []string
	Filenames     []string
	Suffixes      []string
	Shebangs      []string
	LineComments  []string
	BlockComments []*Block
	Quotes        []string
	RawQuotes     []string
	Nested        bool
}
