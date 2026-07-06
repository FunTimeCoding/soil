package main

import "github.com/funtimecoding/soil/pkg/pushover"

func main() {
	p := pushover.NewEnvironment()
	p.Send("test")
}
