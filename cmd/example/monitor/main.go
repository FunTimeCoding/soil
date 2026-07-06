package main

import "github.com/funtimecoding/soil/pkg/monitor/check/collect"

func main() {
	collect.LoopIndividual()

	if false {
		collect.Loop()
		collect.Check(false, false)
	}
}
