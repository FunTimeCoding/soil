package service_tester

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/option"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/service"
	parser "github.com/funtimecoding/soil/pkg/tool/goraidparsed/face"
	"path/filepath"
	"testing"
)

func NewWith(
	t *testing.T,
	logCachePath string,
	outputPath string,
	p parser.Client,
	r face.Reporter,
) *service.Service {
	t.Helper()
	o := option.New()
	o.LitePath = filepath.Join(t.TempDir(), "goraidd.db")
	o.LogCachePath = logCachePath
	o.ElitePath = t.TempDir()
	o.OutputPath = outputPath
	o.Issuer = "https://gate.example.org"
	o.ClientIdentifier = "tester"
	o.ClientSecret = "tester-secret"
	o.EncryptionSecret = "tester-encryption-secret"
	o.PublicLocator = "https://raid.example.org"
	o.Parser = func() parser.Client { return p }
	result := service.New(o, logger.New(t.Context()), r)
	t.Cleanup(result.Close)

	return result
}
