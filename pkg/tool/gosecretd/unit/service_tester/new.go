package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/face"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/option"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/service"
	"testing"
)

func New(
	t *testing.T,
	c *mock_client.Client,
) *service.Service {
	t.Helper()
	o := option.New()
	o.Source = func() face.SecretSource { return c }
	result := service.New(o)
	t.Cleanup(result.Close)

	return result
}
