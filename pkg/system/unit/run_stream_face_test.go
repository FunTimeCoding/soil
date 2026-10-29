package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/system/run"
	"github.com/funtimecoding/soil/pkg/system/run/stream/mock_stream"
	"testing"
)

func TestTheStreamAndItsMockShareTheFace(t *testing.T) {
	s := run.New().Stream("true")
	assert.FatalOnError(t, s.Wait())
	m := mock_stream.New(nil, false, nil)
	defer m.Close()
	faces := []face.Stream{s, m}
	assert.Count(t, 2, faces)
}
