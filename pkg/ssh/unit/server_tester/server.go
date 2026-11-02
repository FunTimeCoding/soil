package server_tester

import (
	"golang.org/x/crypto/ssh"
	"net"
	"sync"
	"testing"
)

type Server struct {
	t             *testing.T
	listener      net.Listener
	configuration *ssh.ServerConfig
	mutex         sync.Mutex
	connections   []net.Conn
	group         sync.WaitGroup
	closed        sync.Once
	Port          int
	Password      string
}
