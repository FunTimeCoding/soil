package lint

import (
	"github.com/funtimecoding/soil/pkg/lint/option"
	"github.com/funtimecoding/soil/pkg/lint/pointer/retarget"
	"github.com/funtimecoding/soil/pkg/lint/types/path_move"
	"github.com/funtimecoding/soil/pkg/lint/types/retargeted_document"
	"github.com/funtimecoding/soil/pkg/system"
)

func RetargetTree(
	root string,
	moves []*path_move.Move,
) []*retargeted_document.Document {
	o := option.New("", false)
	o.Metadata = true
	repository, _ := Walk(root, o)
	r := newResolver(repository, o)
	var result []*retargeted_document.Document

	for _, p := range markdownFiles(repository.Files, o) {
		content := system.ReadFile(repository.Root, p)

		if !retarget.Mentions(p, content, moves) {
			continue
		}

		after, changes := RetargetDocument(r, p, content, moves)

		if len(changes) == 0 {
			continue
		}

		result = append(
			result,
			retargeted_document.New(p, content, after, changes),
		)
	}

	return result
}
