package entry_page

import (
	"github.com/funtimecoding/soil/pkg/web/component/opening_row"
	"maragu.dev/gomponents"
)

func Removal(identifier uint) gomponents.Node {
	return opening_row.New(rowMark(identifier), "", "").Removal()
}
