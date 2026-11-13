package gobreak

import (
	"github.com/funtimecoding/soil/pkg/text/literal_edit/block"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/break_file"
)

func fileBlocks(
	b *break_file.Break,
	file string,
) ([]*block.Block, string) {
	var result []*block.Block

	for _, d := range b.Edits {
		if d.File != file {
			continue
		}

		k := block.NewSearch(len(result)+1, d.Find, d.Replace)

		if e := k.Validate(); e != nil {
			return nil, e.Error()
		}

		result = append(result, k)
	}

	return result, ""
}
