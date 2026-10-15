package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gobuild"
	"github.com/funtimecoding/soil/pkg/tool/gobuild/option"
	"testing"
)

func TestValidateAcceptsNamesAndTargets(t *testing.T) {
	o := option.New()
	o.Names = []string{"gosed", "golint"}
	o.Targets = []string{"linux-amd64"}
	assert.Nil(t, gobuild.Validate(o))
}

func TestValidateRefusesACommaList(t *testing.T) {
	o := option.New()
	o.Names = []string{"gosed,golint"}
	assert.Error(t, gobuild.Validate(o))
}

func TestValidateRefusesAPathAsName(t *testing.T) {
	o := option.New()
	o.Names = []string{"cmd/gosed"}
	assert.Error(t, gobuild.Validate(o))
}

func TestValidateSplitsModuleAndVersion(t *testing.T) {
	o := option.New()
	o.Names = []string{"gosed"}
	o.Module = "github.com/funtimecoding/soil@v0.11.190"
	assert.Nil(t, gobuild.Validate(o))
	assert.String(t, "github.com/funtimecoding/soil", o.Module)
	assert.String(t, "v0.11.190", o.Version)
}

func TestValidateLeavesAModuleWithoutVersionToLatest(t *testing.T) {
	o := option.New()
	o.Names = []string{"gosed"}
	o.Module = "github.com/funtimecoding/soil"
	assert.Nil(t, gobuild.Validate(o))
	assert.String(t, "", o.Version)
}

func TestValidateRefusesAModuleWithoutNames(t *testing.T) {
	o := option.New()
	o.Module = "github.com/funtimecoding/soil"
	assert.Error(t, gobuild.Validate(o))
}

func TestValidateRefusesAnEmptyVersion(t *testing.T) {
	o := option.New()
	o.Names = []string{"gosed"}
	o.Module = "github.com/funtimecoding/soil@"
	assert.Error(t, gobuild.Validate(o))
}

func TestValidateRefusesMainWithAModule(t *testing.T) {
	o := option.New()
	o.Names = []string{"gosed"}
	o.Module = "github.com/funtimecoding/soil"
	o.Main = "cmd/gosed/main.go"
	assert.Error(t, gobuild.Validate(o))
}

func TestValidateRefusesMainWithSeveralNames(t *testing.T) {
	o := option.New()
	o.Names = []string{"gosed", "golint"}
	o.Main = "cmd/gosed/main.go"
	assert.Error(t, gobuild.Validate(o))
}

func TestValidateRefusesOutputWithSeveralTargets(t *testing.T) {
	o := option.New()
	o.Names = []string{"gosed"}
	o.Targets = []string{"linux-amd64", "darwin-arm64"}
	o.Output = "tmp/gosed"
	assert.Error(t, gobuild.Validate(o))
}
