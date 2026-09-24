package integration

import "github.com/funtimecoding/soil/pkg/git/unit/repository_tester"

func commit(
	directory string,
	message string,
) {
	repository_tester.NewAt(directory).Commit(message)
}
