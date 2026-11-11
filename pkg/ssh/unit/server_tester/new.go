package server_tester

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
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
	curve, e := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.FatalOnError(t, e)
	second, e := ssh.NewSignerFromKey(curve)
	assert.FatalOnError(t, e)
	l, e := net.Listen("tcp", "127.0.0.1:0")
	assert.FatalOnError(t, e)
	result := &Server{
		t:        t,
		listener: l,
		Port:     l.Addr().(*net.TCPAddr).Port,
		Password: "secret",
		HostKey:  signer.PublicKey(),
	}
	result.configuration = &ssh.ServerConfig{
		PasswordCallback: result.checkPassword,
	}
	result.configuration.AddHostKey(signer)
	result.configuration.AddHostKey(second)
	result.group.Add(1)
	go result.serve()
	t.Cleanup(result.Close)

	return result
}
