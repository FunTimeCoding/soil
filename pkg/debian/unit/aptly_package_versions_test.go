package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/debian/aptly"
	"testing"
)

func TestPackageVersionsReadsTheAptlyKeyFormat(t *testing.T) {
	v := aptly.PackageVersions(packageKeys(), "foxtrot")
	assert.Count(t, 3, v)
	assert.String(t, "1.2.1", v[0])
	assert.String(t, "1.2.5", v[1])
}

func TestPackageVersionsRefusesAnotherPackage(t *testing.T) {
	assert.Count(t, 2, aptly.PackageVersions(packageKeys(), "gohw"))
}

func TestPackageVersionsRefusesAnAbsentPackage(t *testing.T) {
	assert.Count(t, 0, aptly.PackageVersions(packageKeys(), "gooutpostd"))
}
