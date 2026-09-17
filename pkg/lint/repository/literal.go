package repository

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func (r *Repository) Literal(
	directory string,
	needle string,
) bool {
	contents, found := r.contents[directory]

	if !found {
		if strings.HasPrefix(directory, constant.ParentDirectory) {
			contents = siblingContents(r.Absolute(directory))
		} else {
			prefix := join.Empty(directory, stringsConstant.Slash)

			for _, p := range r.Files.Files() {
				if strings.HasSuffix(p, constant.MarkdownExtension) {
					continue
				}

				if directory == "" || strings.HasPrefix(p, prefix) {
					contents = append(contents, r.Files.ReadString(p))
				}
			}
		}

		r.contents[directory] = contents
	}

	for _, content := range contents {
		if pointer.ContainsLiteral(content, needle) {
			return true
		}
	}

	return false
}
