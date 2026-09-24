package classifier

import "github.com/funtimecoding/soil/pkg/measure/language"

type Classifier struct {
	language *language.Language
	open     []*language.Block
	quote    string
	raw      bool
}
