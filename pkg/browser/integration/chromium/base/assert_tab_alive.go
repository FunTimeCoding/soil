package base

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"time"
)

func (s *Stack) AssertTabAlive(identifier string) {
	s.T.Helper()
	time.Sleep(2 * time.Second)
	assert.True(s.T, s.TabAlive(identifier))
}
