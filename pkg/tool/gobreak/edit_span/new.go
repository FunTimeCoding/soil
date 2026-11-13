package edit_span

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/text/literal_edit/match"
	"path"
	"strings"
)

func New(
	modulePath string,
	file string,
	content string,
	m *match.Match,
) *Span {
	matched := strings.TrimSuffix(
		content[m.Offset:m.Offset+m.Length],
		stringsConstant.Unix,
	)

	return &Span{
		Profile: path.Join(modulePath, file),
		File:    file,
		First:   m.Line,
		Last:    m.Line + strings.Count(matched, stringsConstant.Unix),
	}
}
