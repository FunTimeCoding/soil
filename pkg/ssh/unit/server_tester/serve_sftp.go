package server_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func serveSftp(c ssh.Channel) {
	v, e := sftp.NewServer(c)

	if e != nil {
		release(c)

		return
	}

	defer release(v)
	errors.LogOnError(v.Serve())
}
