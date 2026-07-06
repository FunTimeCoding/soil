package main

import "github.com/funtimecoding/soil/pkg/errors/sentry/example"

func main() {
	example.Whoami()

	if false {
		example.Capture()
		example.Issue()
		example.TrackedIssue()
		example.Organization()
		example.Project()
	}
}
