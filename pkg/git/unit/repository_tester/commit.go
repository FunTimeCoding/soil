package repository_tester

import "github.com/funtimecoding/soil/pkg/constant"

func (r *Tester) Commit(message string) {
	Command(r.Clone, "add", constant.CurrentDirectory)
	Command(r.Clone, "commit", "-m", message)
}
