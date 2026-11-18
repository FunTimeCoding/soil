package fixture

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/nut/constant"
	"net"
)

func NewServer(answers map[string]string) *Server {
	l, e := net.Listen(constant.Network, "127.0.0.1:0")
	errors.PanicOnError(e)
	result := &Server{listener: l, answers: answers}
	go result.serve()

	return result
}
