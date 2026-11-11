package lint

import (
	"github.com/funtimecoding/soil/pkg/lint/option"
	"github.com/funtimecoding/soil/pkg/lint/types/move_set"
	"github.com/funtimecoding/soil/pkg/lint/types/retargeted_document"
)

func RetargetTree(
	root string,
	moves *move_set.Set,
) []*retargeted_document.Document {
	o := option.New("", false)
	o.Metadata = true
	repository, _ := Walk(root, o)

	return RetargetFiles(
		repository,
		o,
		markdownFiles(repository.Files, o),
		moves,
	)
}
