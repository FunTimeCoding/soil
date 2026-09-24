package integration

import "github.com/funtimecoding/soil/pkg/git/unit/repository_tester"

func command(
	directory string,
	arguments ...string,
) string {
	return repository_tester.Command(directory, arguments...)
}
