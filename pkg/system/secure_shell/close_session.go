package secure_shell

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"golang.org/x/crypto/ssh"
)

func CloseSession(s *ssh.Session) {
	if e := s.Close(); e != nil && e.Error() != constant.EndOfFile {
		errors.LogOnError(e)
	}
}
