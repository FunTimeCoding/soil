package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	go_mod "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/go_mod/dependency"
	"github.com/funtimecoding/soil/pkg/go_mod/license"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"testing"
)

func TestLicenseScanPermissive(t *testing.T) {
	assert.Strings(
		t,
		[]string{"MIT"},
		license.Scan(fixture.Read(constant.LicensePath, "mit.txt")),
	)
	assert.Strings(
		t,
		[]string{"Apache-2.0"},
		license.Scan(fixture.Read(constant.LicensePath, "apache.txt")),
	)
}

func TestLicenseScanCopyleft(t *testing.T) {
	assert.Strings(
		t,
		[]string{"MPL-2.0"},
		license.Scan(fixture.Read(constant.LicensePath, "mozilla.txt")),
	)
	assert.Strings(
		t,
		[]string{"GPL-3.0-or-later"},
		license.Scan(fixture.Read(constant.LicensePath, "gnu.txt")),
	)
}

func TestLicenseScanFiller(t *testing.T) {
	assert.Count(
		t,
		0,
		license.Scan(fixture.Read(constant.LicensePath, "filler.txt")),
	)
}

func TestLicenseFamily(t *testing.T) {
	assert.String(t, "permissive", license.Family([]string{"MIT"}))
	assert.String(
		t,
		"permissive",
		license.Family([]string{"BSD-3-Clause", "Apache-2.0"}),
	)
	assert.String(t, "weak_copyleft", license.Family([]string{"MPL-2.0"}))
	assert.String(
		t,
		"weak_copyleft",
		license.Family([]string{"MIT", "LGPL-2.1"}),
	)
	assert.String(
		t,
		"copyleft",
		license.Family([]string{"MIT", "GPL-3.0-or-later"}),
	)
	assert.String(t, "copyleft", license.Family([]string{"AGPL-3.0"}))
	assert.String(t, "license_unknown", license.Family([]string{"WTFPL"}))
	assert.String(t, "license_unknown", license.Family(nil))
}

func TestLicenseFindFiles(t *testing.T) {
	assert.Strings(
		t,
		[]string{"LICENSE"},
		license.FindFiles(
			fixture.Path(constant.LicensePath, "module", "with_license"),
		),
	)
	assert.Count(
		t,
		0,
		license.FindFiles(
			fixture.Path(constant.LicensePath, "module", "without_license"),
		),
	)
	assert.Strings(
		t,
		[]string{"LICENSE.BSD", "LICENSE.MPL"},
		license.FindFiles(
			fixture.Path(constant.LicensePath, "module", "dual_license"),
		),
	)
}

func TestDependencyValidateDualLicense(t *testing.T) {
	d := dependency.New(
		"host.example/charlie",
		"v1.0.0",
		fixture.Path(constant.LicensePath, "module", "dual_license"),
	)
	d.Validate()
	assert.Strings(t, []string{"BSD-3-Clause", "MPL-2.0"}, d.Identifiers)
	assert.String(t, "weak_copyleft", d.Family)
	assert.Strings(t, []string{"weak_copyleft"}, d.Concerns())
}

func TestDependencyValidateWithLicense(t *testing.T) {
	d := dependency.New(
		"host.example/alpha",
		"v1.0.0",
		fixture.Path(constant.LicensePath, "module", "with_license"),
	)
	d.Validate()
	assert.Strings(t, []string{"LICENSE"}, d.Files)
	assert.Strings(t, []string{"MIT"}, d.Identifiers)
	assert.String(t, "permissive", d.Family)
	assert.False(t, d.HasConcerns())
}

func TestDependencyValidateWithoutLicense(t *testing.T) {
	d := dependency.New(
		"host.example/bravo",
		"v1.0.0",
		fixture.Path(constant.LicensePath, "module", "without_license"),
	)
	d.Validate()
	assert.Count(t, 0, d.Files)
	assert.String(t, "license_unknown", d.Family)
	assert.Strings(t, []string{"license_missing"}, d.Concerns())
	assert.True(t, d.HasConcern(go_mod.LicenseMissing))
}
