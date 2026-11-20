package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/store_tester"
	"testing"
	"time"
)

func TestSearchPlacementsMatchesPartOfName(t *testing.T) {
	o := store_tester.New(t)
	now := time.Now()
	o.Save(constant.SourceKubernetes, "gitlab-runner", now)
	o.Save(constant.SourceKubernetes, "gitlab", now)
	o.Save(constant.SourceProcess, "goqueryd", now)
	v, e := o.Store.SearchPlacements("gitlab")
	assert.FatalOnError(t, e)
	assert.Count(t, 2, v)
	assert.String(t, "gitlab", v[0].Name)
	assert.String(t, "gitlab-runner", v[1].Name)
}

func TestSearchPlacementsWithoutNameCarriesEverything(t *testing.T) {
	o := store_tester.New(t)
	now := time.Now()
	o.Save(constant.SourceKubernetes, "gitlab", now)
	o.Save(constant.SourceProcess, "goqueryd", now)
	v, e := o.Store.SearchPlacements("")
	assert.FatalOnError(t, e)
	assert.Count(t, 2, v)
}

func TestSearchPlacementsRefusesAbsentName(t *testing.T) {
	o := store_tester.New(t)
	o.Save(constant.SourceKubernetes, "gitlab", time.Now())
	v, e := o.Store.SearchPlacements("nowhere")
	assert.FatalOnError(t, e)
	assert.Count(t, 0, v)
}
