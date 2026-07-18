package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"testing"
)

func TestStoreRoundTrip(t *testing.T) {
	s := New(lite.NewMemory())
	s.MustSaveChannel(-100, "announcements")
	s.MustSaveChannel(-100, "announcements")
	s.MustSaveUser(7, "admin")
	assert.Integer(t, 1, len(s.MustChannels()))
	assert.Integer(t, 1, len(s.MustUsers()))
	h := s.MustChannelByName("announcements")
	assert.NotNil(t, h)
	assert.Integer64(t, -100, h.Identifier)
	assert.True(t, s.MustChannelByName("missing") == nil)
}
