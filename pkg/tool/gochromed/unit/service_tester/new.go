package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gochromed/face"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/option"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/service"
	"testing"
)

func New(
	t *testing.T,
	c *mock_client.Client,
) *service.Service {
	t.Helper()
	o := option.New()
	o.DownloadDirectory = t.TempDir()
	o.Source = func() face.ChromiumSource { return c }
	result := service.New(o)
	t.Cleanup(result.Close)

	return result
}
