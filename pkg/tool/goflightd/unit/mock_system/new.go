package mock_system

import "testing"

func New(t *testing.T) *System {
	t.Helper()
	result := &System{
		outputs:  make(map[string]string),
		failures: make(map[string]error),
	}
	t.Cleanup(result.Close)

	return result
}
