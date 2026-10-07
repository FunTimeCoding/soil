package language

import "github.com/funtimecoding/soil/pkg/measure/types/language_block"

func (l *Language) WithBlockComment(
	open string,
	close string,
) *Language {
	l.BlockComments = append(l.BlockComments, language_block.New(open, close))

	return l
}
