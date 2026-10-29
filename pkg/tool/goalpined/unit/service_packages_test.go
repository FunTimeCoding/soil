package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/unit/index_tester"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/unit/service_tester"
	"path/filepath"
	"testing"
)

func TestPackagesReadsTheIndexesUnderTheRoot(t *testing.T) {
	o := service_tester.New(t)
	index_tester.WriteIndex(
		filepath.Join(o.PackageRoot, "v3.21", "main", "x86_64"),
	)
	all, e := o.Service.Packages("")
	assert.FatalOnError(t, e)
	assert.Count(t, 1, all)
	assert.String(t, "v3.21", all[0].Version)
	assert.Count(t, 2, all[0].Packages)
	one, f := o.Service.Packages("gohw")
	assert.FatalOnError(t, f)
	assert.Count(t, 1, one[0].Packages)
}
