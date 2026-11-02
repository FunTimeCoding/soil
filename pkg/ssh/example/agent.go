package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/ssh"
	"github.com/funtimecoding/soil/pkg/system"
)

func Agent(host string) {
	s := ssh.New(system.User().Username, host, true)

	if e := s.Dial(); e != nil {
		console.Format("Dial: %v\n", e)

		return
	}

	defer s.Close()
	console.Format("Hostname: %s\n", s.Run("hostname").OutputString)
	console.Format("Uptime: %s\n", s.Run("uptime").OutputString)
	entries, e := s.ListDirectory(constant.CurrentDirectory)
	errors.PanicOnError(e)
	console.Format("Home: %d entries\n", len(entries))
}
