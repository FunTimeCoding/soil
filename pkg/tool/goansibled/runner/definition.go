package runner

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/markup"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/constant"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/types/host_command"
	"io/fs"
	"os"
	"path/filepath"
)

func (r *Runner) definition(name string) (*host_command.HostCommand, error) {
	b, e := os.ReadFile(
		filepath.Join(
			r.cloneDirectory,
			r.ansibleDirectory,
			constant.HostCommandFile,
		),
	)

	if errors.Is(e, fs.ErrNotExist) {
		return nil, not_found.New(constant.HostCommandKind, name)
	}

	if e != nil {
		return nil, e
	}

	definitions := map[string]*host_command.HostCommand{}

	if f := markup.Decode(string(b), &definitions); f != nil {
		return nil, f
	}

	result, found := definitions[name]

	if !found {
		return nil, not_found.New(constant.HostCommandKind, name)
	}

	return result, nil
}
