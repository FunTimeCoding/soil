package fixture

import (
	"bufio"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/nut/constant"
	"io"
	"net"
)

func (s *Server) answer(n net.Conn) {
	defer errors.LogClose(n)
	r := bufio.NewScanner(n)

	for r.Scan() {
		command := r.Text()

		if command == constant.Logout {
			_, e := io.WriteString(n, constant.FixtureGoodbye)
			errors.LogOnError(e)

			return
		}

		a, found := s.answers[command]

		if !found {
			a = constant.FixtureUnknownCommand
		}

		_, e := io.WriteString(n, a)
		errors.LogOnError(e)
	}
}
