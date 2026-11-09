package mock_dump

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"
	"os"
)

func (p *Program) Dump(
	_ context.Context,
	i *instance.Instance,
	path string,
) error {
	p.Dumped = append(p.Dumped, i.Name)

	if e := os.WriteFile(
		path,
		[]byte(i.Database),
		constant.BackupFileMode,
	); e != nil {
		return e
	}

	return p.Failure
}
