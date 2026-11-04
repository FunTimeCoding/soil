package salt_tester

import "github.com/funtimecoding/soil/pkg/provision/constant"

func (s *Server) HoldKeys() {
	s.holdPath(constant.SaltKeysPath)
}
