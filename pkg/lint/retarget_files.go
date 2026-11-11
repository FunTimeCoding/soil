package lint

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/lint/option"
	"github.com/funtimecoding/soil/pkg/lint/pointer/retarget"
	"github.com/funtimecoding/soil/pkg/lint/repository"
	"github.com/funtimecoding/soil/pkg/lint/types/move_set"
	"github.com/funtimecoding/soil/pkg/lint/types/retargeted_document"
	"github.com/funtimecoding/soil/pkg/system"
	"io/fs"
)

func RetargetFiles(
	p *repository.Repository,
	o *option.Lint,
	paths []string,
	moves *move_set.Set,
) []*retargeted_document.Document {
	r := newResolver(p, o)
	root := system.Root(p.Root)
	defer errors.PanicClose(root)
	var result []*retargeted_document.Document

	for _, path := range paths {
		b, e := root.ReadFile(path)

		if errors.Is(e, fs.ErrNotExist) {
			continue
		}

		errors.PanicOnError(e)
		content := string(b)

		if !retarget.Mentions(path, content, moves) {
			continue
		}

		after, changes := RetargetDocument(r, path, content, moves)

		if len(changes) == 0 {
			continue
		}

		result = append(
			result,
			retargeted_document.New(path, content, after, changes),
		)
	}

	return result
}
