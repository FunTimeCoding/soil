package ssh

import (
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
)

type Client struct {
	user         string
	host         string
	secure       bool
	authenticate func() (ssh.AuthMethod, io.Closer, error)
	client       *ssh.Client
	sftp         *sftp.Client
	Panic        bool
}
