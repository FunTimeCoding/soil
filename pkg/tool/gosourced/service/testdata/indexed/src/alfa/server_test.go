package alfa

import "testing"

func TestServer(t *testing.T) {
	NewServer().Start()
	restart()
}
