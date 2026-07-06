package main

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/check/job"
	"github.com/funtimecoding/soil/pkg/kubernetes/example"
)

func main() {
	job.Run()

	if false {
		example.Event()
		example.Pod()
		example.Namespace()
		example.Node()
	}
}
