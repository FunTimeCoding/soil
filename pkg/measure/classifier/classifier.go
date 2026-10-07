package classifier

import (
	"github.com/funtimecoding/soil/pkg/measure/language"
	"github.com/funtimecoding/soil/pkg/measure/types/language_block"
)

type Classifier struct {
	language *language.Language
	open     []*language_block.Block
	quote    string
	raw      bool
}
