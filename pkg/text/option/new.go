package option

func New() *Whitespace {
	return &Whitespace{NewlineAtEnd: true, AllowedBlankLines: 1}
}
