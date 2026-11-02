package face

import (
	"github.com/funtimecoding/soil/pkg/ssh/command"
	"github.com/funtimecoding/soil/pkg/system/result"
	"io/fs"
)

type SSHClient interface {
	Dial() error
	Run(s string) *result.Result
	RunCommand(o *command.Command) *result.Result
	WriteFile(
		path string,
		content []byte,
	) error
	ReadFile(path string) ([]byte, error)
	ListDirectory(path string) ([]fs.FileInfo, error)
	RemoveFile(path string) error
	Close()
}
