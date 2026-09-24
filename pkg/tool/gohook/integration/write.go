package integration

import "github.com/funtimecoding/soil/pkg/git/unit/repository_tester"

func write(
	root string,
	name string,
	content string,
) {
	repository_tester.NewAt(root).Write(name, content)
}
