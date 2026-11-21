package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/daemon"
	debianMock "github.com/funtimecoding/soil/pkg/vulnerability/debian/mock_client"
	nvdMock "github.com/funtimecoding/soil/pkg/vulnerability/national_database/mock_client"
	"testing"
)

type Server struct {
	NationalDatabase *nvdMock.Client
	SecurityTracker  *debianMock.Client
	ExtendedTracker  *debianMock.Client
	Reporter         *memory.Memory
	daemon           *daemon.Daemon
	served           *model_context_server.Server
	t                *testing.T
}
