package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/provision/salt"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/option"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/service"
	"path/filepath"
	"testing"
)

func TestServiceCloseReleasesTheDatabase(t *testing.T) {
	o := option.New()
	o.LitePath = filepath.Join(t.TempDir(), "gosaltd.db")
	o.Master = func() *salt.Client {
		return salt.New("localhost", 1, "guard", "guard")
	}
	s := service.New(o, logger.New(t.Context()), memory.New())
	s.Close()
	_, e := s.Recent(1)
	assert.True(t, e != nil)
}
