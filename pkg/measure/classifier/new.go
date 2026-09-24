package classifier

import "github.com/funtimecoding/soil/pkg/measure/language"

func New(l *language.Language) *Classifier {
	return &Classifier{language: l}
}
