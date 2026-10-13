package main

import "github.com/funtimecoding/soil/pkg/pushover"

func main() {
	p := pushover.NewEnvironment()
	defer p.Close()
	p.Send("test")
}
