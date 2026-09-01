package main

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/linux/systemd"
	"github.com/funtimecoding/soil/pkg/ssh"
	"github.com/funtimecoding/soil/pkg/ssh/command"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/time"
)

func main() {
	a := argument.NewSimple("systemd-example")
	a.ParseSimple()
	host := a.RequiredPositional(0, "HOST")
	service := a.RequiredPositional(1, "SERVICE")
	console.Format("Host: %s\n", host)
	console.Format("Service: %s\n", service)
	s := ssh.New(system.User().Username, host, true)
	defer s.Close()
	c := systemd.New(s)

	if false {
		console.Format("List: %+v\n", c.ListOutput())
		console.Format("Not found: %+v\n", c.NotFoundOutput())
		console.Format("Status: %+v\n", c.Status(service))
	}

	if false {
		console.Format("Is active running: %v\n", c.IsActiveRunning(service))
		console.Format("Start time: %v\n", time.Format(c.StartTime(service)))
	}

	if true {
		s.RunCommand(command.New("sudo echo $TEST").Set("TEST", "yay")).Print()
	}
}
