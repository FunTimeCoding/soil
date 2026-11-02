package ssh

import (
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"sync"
)

type Client struct {
	user         string
	host         string
	port         int
	secure       bool
	authenticate func() (ssh.AuthMethod, io.Closer, error)
	client       *ssh.Client
	sftp         *sftp.Client
	mutex        sync.Mutex
	Panic        bool
}
