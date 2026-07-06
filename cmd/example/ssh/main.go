package main

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/ssh"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/join"
	"github.com/funtimecoding/soil/pkg/system/secure_shell"
)

func main() {
	a := argument.NewSimple("ssh-example")
	a.ParseSimple()
	n := a.RequiredPositional(0, "NODE")
	fmt.Printf("Node: %s\n", n)

	if false {
		s := ssh.New(system.User().Username, n, false)
		defer s.Close()
		r := s.Run("ls")
		fmt.Printf("Run: %s\n", r.OutputString)
	}

	if true {
		s := ssh.NewWithFile(
			system.User().Username,
			n,
			join.Absolute(
				system.Home(),
				secure_shell.ConfigurationDirectory,
				"ansible",
			),
			"id_rsa_insecure",
			false,
		)
		defer s.Close()
		r := s.Run("ls")
		fmt.Printf("Run: %s\n", r.OutputString)
	}
}
