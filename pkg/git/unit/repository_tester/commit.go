package repository_tester

import (
	"github.com/funtimecoding/soil/pkg/constant"
	gitConstant "github.com/funtimecoding/soil/pkg/git/constant"
)

func (r *Tester) Commit(message string) {
	Command(r.Clone, "add", constant.CurrentDirectory)
	Command(r.Clone, gitConstant.Commit, gitConstant.MessageArgument, message)
}
