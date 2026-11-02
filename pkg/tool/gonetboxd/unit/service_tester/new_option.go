package service_tester

import (
	"github.com/funtimecoding/soil/pkg/netbox/face"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/option"
	"github.com/prometheus/client_golang/prometheus"
	"path/filepath"
	"testing"
)

func NewOption(
	t *testing.T,
	c face.NetBoxClient,
) *option.Netbox {
	t.Helper()
	result := option.New()
	result.LitePath = filepath.Join(t.TempDir(), "gonetboxd.db")
	result.Source = func(prometheus.ObserverVec) face.NetBoxClient { return c }

	return result
}
