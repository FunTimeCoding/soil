package gobreak

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/text/literal_edit/match"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/break_file"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/edit_span"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/overlay"
	"os"
	"path/filepath"
)

func applyFile(
	root string,
	modulePath string,
	file string,
	b *break_file.Break,
	o *overlay.Overlay,
) ([]*edit_span.Span, string) {
	path := filepath.Join(root, file)
	content, e := os.ReadFile(path)
	errors.PanicOnError(e)
	blocks, failure := fileBlocks(b, file)

	if failure != "" {
		return nil, fmt.Sprintf(constant.UnmatchedFormat, file, failure)
	}

	matches, failures := match.Locate(string(content), blocks)

	if len(failures) > 0 {
		return nil, fmt.Sprintf(
			constant.UnmatchedFormat,
			file,
			join.SemicolonSpace(failures),
		)
	}

	var result []*edit_span.Span

	for _, m := range matches {
		result = append(
			result,
			edit_span.New(modulePath, file, string(content), m),
		)
	}

	o.Add(path, match.Apply(string(content), matches))

	return result, ""
}
