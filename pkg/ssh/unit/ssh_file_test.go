package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/dropped"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/ssh"
	"github.com/funtimecoding/soil/pkg/ssh/unit/server_tester"
	"os"
	"path/filepath"
	"sync"
	"testing"
)

func TestAFileRoundTripsOverSftp(t *testing.T) {
	c := server_tester.New(t).Dial()
	directory := t.TempDir()
	path := filepath.Join(directory, "alfa.yaml")
	assert.FatalOnError(t, c.WriteFile(path, []byte("#cloud-config")))
	b, e := c.ReadFile(path)
	assert.FatalOnError(t, e)
	assert.String(t, "#cloud-config", string(b))
	entries, e := c.ListDirectory(directory)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, entries)
	assert.String(t, "alfa.yaml", entries[0].Name())
	assert.FatalOnError(t, c.RemoveFile(path))
	empty, e := c.ListDirectory(directory)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, empty)
}

func TestAMissingFileIsNotFound(t *testing.T) {
	c := server_tester.New(t).Dial()
	path := filepath.Join(t.TempDir(), "alfa.yaml")
	_, e := c.ReadFile(path)
	assert.True(t, not_found.Is(e))
	assert.True(t, not_found.Is(c.RemoveFile(path)))
	_, f := c.ListDirectory(path)
	assert.True(t, not_found.Is(f))
}

func TestARefusedWriteIsForbidden(t *testing.T) {
	c := server_tester.New(t).Dial()
	directory := t.TempDir()
	assert.FatalOnError(t, os.Chmod(directory, 0500))
	t.Cleanup(func() { assert.FatalOnError(t, os.Chmod(directory, 0700)) })
	e := c.WriteFile(filepath.Join(directory, "alfa.yaml"), []byte("bravo"))
	assert.True(t, forbidden.Is(e))
}

func TestALostConnectionIsDropped(t *testing.T) {
	s := server_tester.New(t)
	c := s.Dial()
	path := filepath.Join(t.TempDir(), "alfa.yaml")
	assert.FatalOnError(t, c.WriteFile(path, []byte("bravo")))
	s.Close()
	_, e := c.ReadFile(path)
	assert.True(t, dropped.Is(e))
}

func TestAnUndialledClientAnswersAnError(t *testing.T) {
	c := ssh.NewWithPassword("alfa", "localhost", "", false)
	defer c.Close()
	_, e := c.ReadFile("/alfa")
	assert.NotNil(t, e)
}

func TestConcurrentCallsShareOneSftpSession(t *testing.T) {
	c := server_tester.New(t).Dial()
	path := filepath.Join(t.TempDir(), "alfa.yaml")
	assert.FatalOnError(t, os.WriteFile(path, []byte("bravo"), 0600))
	var group sync.WaitGroup
	failures := make(chan error, 8)

	for range 8 {
		group.Go(
			func() {
				_, e := c.ReadFile(path)
				failures <- e
			},
		)
	}

	group.Wait()
	close(failures)

	for e := range failures {
		assert.FatalOnError(t, e)
	}
}
