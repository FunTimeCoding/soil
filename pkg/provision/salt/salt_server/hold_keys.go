package salt_server

import "github.com/funtimecoding/soil/pkg/provision/constant"

func (s *Server) HoldKeys() {
	s.holdPath(constant.SaltKeysPath)
}
