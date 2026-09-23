package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"testing"
)

func TestFrontendsReadsLayout(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/tool/gotestd/web/new.go",
		constant.FixtureFrontendNewSample,
	)
	v.WriteString(
		"pkg/tool/gotestd/web/mount.go",
		constant.FixtureFrontendMountSample,
	)
	v.WriteString("pkg/tool/gotestd/web/favicon.png", "png")
	f := frontends(v)
	assert.Integer(t, 1, len(f))
	assert.String(t, "Straw", f[0].Theme)
	assert.Boolean(t, true, f[0].Style)
	assert.Boolean(t, true, f[0].Palette)
	assert.Boolean(t, true, f[0].PaletteRoute)
	assert.Integer(t, 2, f[0].Items)
	assert.Boolean(t, false, f[0].Live)
	assert.Boolean(t, true, f[0].Favicon)
	assert.Boolean(t, true, f[0].FaviconRoute)
}

func TestFrontendsSkipsNonLayoutWebPackage(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/tool/gotestd/web/new.go",
		"package web\n\nfunc Get() {}\n",
	)
	assert.Integer(t, 0, len(frontends(v)))
}

func TestFrontendsReportsMissingRoutes(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/tool/gotestd/web/new.go",
		constant.FixtureFrontendNewSample,
	)
	f := frontends(v)
	assert.Integer(t, 1, len(f))
	assert.Boolean(t, true, f[0].Palette)
	assert.Boolean(t, false, f[0].PaletteRoute)
	assert.Boolean(t, false, f[0].Favicon)
	assert.Boolean(t, false, f[0].FaviconRoute)
}
