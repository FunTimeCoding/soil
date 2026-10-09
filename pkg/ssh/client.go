package ssh

import (
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type Client struct {
	user         string
	host         string
	secure       bool
	authenticate func() (ssh.AuthMethod, error)
	client       *ssh.Client
	sftp         *sftp.Client
	Panic        bool
}
