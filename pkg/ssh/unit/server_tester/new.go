package server_tester

import (
	"crypto/ed25519"
	"crypto/rand"
	"github.com/funtimecoding/soil/pkg/assert"
	"golang.org/x/crypto/ssh"
	"net"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	_, key, e := ed25519.GenerateKey(rand.Reader)
	assert.FatalOnError(t, e)
	signer, e := ssh.NewSignerFromKey(key)
	assert.FatalOnError(t, e)
	l, e := net.Listen("tcp", "127.0.0.1:0")
	assert.FatalOnError(t, e)
	result := &Server{
		t:        t,
		listener: l,
		Port:     l.Addr().(*net.TCPAddr).Port,
		Password: "secret",
	}
	result.configuration = &ssh.ServerConfig{
		PasswordCallback: result.checkPassword,
	}
	result.configuration.AddHostKey(signer)
	result.group.Add(1)
	go result.serve()
	t.Cleanup(result.Close)

	return result
}
