package main

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/ssh"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
)

func main() {
	a := argument.NewSimple("ssh-example")
	a.ParseSimple()
	n := a.RequiredPositional(0, "NODE")
	console.Format("Node: %s\n", n)

	if false {
		s := ssh.New(system.User().Username, n, false)
		defer s.Close()
		r := s.Run("ls")
		console.Format("Run: %s\n", r.OutputString)
	}

	if true {
		s := ssh.NewWithFile(
			system.User().Username,
			n,
			join.Absolute(
				system.Home(),
				constant.SecureShellConfigurationDirectory,
				"ansible",
			),
			"id_rsa_insecure",
			false,
		)
		defer s.Close()
		r := s.Run("ls")
		console.Format("Run: %s\n", r.OutputString)
	}
}
