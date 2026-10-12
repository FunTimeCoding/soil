package repository_tester

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/git/constant"
)

func (r *Tester) Commit(message string) {
	Command(r.Clone, "add", soil.CurrentDirectory)
	Command(r.Clone, constant.Commit, constant.MessageArgument, message)
}
