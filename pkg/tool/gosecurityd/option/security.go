package option

import (
	provision "github.com/funtimecoding/soil/pkg/provision/face"
	"github.com/funtimecoding/soil/pkg/vulnerability/face"
	"github.com/funtimecoding/soil/pkg/web/option"
	"time"
)

type Security struct {
	option.Web
	PostgresLocator   string
	LitePath          string
	CacheDirectory    string
	StaleAfter        time.Duration
	TrackerStaleAfter time.Duration
	Clock             func() time.Time
	NationalDatabase  func() face.NationalDatabase
	SecurityTracker   func() face.SecurityTracker
	ExtendedTracker   func() face.SecurityTracker
	Ansible           func() provision.Client
}
