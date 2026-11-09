package dump

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"strings"
)

func (p *Program) Count(
	x context.Context,
	path string,
) (int, error) {
	listed, e := p.run(
		x,
		nil,
		constant.Identity.Name(),
		constant.RestoreProgram,
		"--list",
		path,
	)

	if e != nil {
		return 0, e
	}

	return strings.Count(listed, constant.TableContentMarker), nil
}
