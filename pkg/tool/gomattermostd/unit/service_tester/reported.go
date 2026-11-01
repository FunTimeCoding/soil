package service_tester

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/face"
	soilFace "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	goclauded "github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/option"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/service"
	"path/filepath"
	"testing"
)

func Reported(
	t *testing.T,
	c face.Client,
	k goclauded.Client,
	configure func(*option.Mattermost),
	r soilFace.Reporter,
) *service.Service {
	t.Helper()
	o := option.New()
	o.LitePath = filepath.Join(t.TempDir(), "mattermost.db")
	o.Mattermost = func() face.Client { return c }
	o.Connector = func() goclauded.Client { return k }

	if configure != nil {
		configure(o)
	}

	result := service.New(o, logger.New(t.Context()), r)
	t.Cleanup(result.Close)

	return result
}
