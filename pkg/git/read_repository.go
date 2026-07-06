package git

import (
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/git/repository"
	"github.com/funtimecoding/soil/pkg/system/run"
)

func ReadRepository(path string) *repository.Repository {
	r := run.New()
	r.Directory = path
	r.Start(constant.Command, constant.RevParse, constant.GitDirectory)
	r.Start(constant.Command, constant.Status, constant.Porcelain)

	return repository.New(path, r.OutputString)
}
