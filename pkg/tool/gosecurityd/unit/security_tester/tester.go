package security_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/debian"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/flag"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/score"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/vulnerability"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store"
	debianMock "github.com/funtimecoding/soil/pkg/vulnerability/debian/mock_client"
	nvdMock "github.com/funtimecoding/soil/pkg/vulnerability/national_database/mock_client"
	"testing"
	"time"
)

type Tester struct {
	NationalDatabase *nvdMock.Client
	SecurityTracker  *debianMock.Client
	ExtendedTracker  *debianMock.Client
	Cache            string
	Later            time.Duration
	Reporter         *memory.Memory
	vulnerability    *vulnerability.Service
	score            *score.Service
	debian           *debian.Service
	flag             *flag.Service
	store            *store.Store
	t                *testing.T
}
