package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/store_tester"
	"testing"
	"time"
)

func TestSearchPlacementsMatchesPartOfTheName(t *testing.T) {
	s := store_tester.NewStore(t)
	now := time.Now()
	store_tester.Save(t, s, constant.SourceKubernetes, "gitlab-runner", now)
	store_tester.Save(t, s, constant.SourceKubernetes, "gitlab", now)
	store_tester.Save(t, s, constant.SourceProcess, "goqueryd", now)
	v, e := s.SearchPlacements("gitlab")
	assert.FatalOnError(t, e)
	assert.Count(t, 2, v)
	assert.String(t, "gitlab", v[0].Name)
	assert.String(t, "gitlab-runner", v[1].Name)
}

func TestSearchPlacementsWithoutANameCarriesEverything(t *testing.T) {
	s := store_tester.NewStore(t)
	now := time.Now()
	store_tester.Save(t, s, constant.SourceKubernetes, "gitlab", now)
	store_tester.Save(t, s, constant.SourceProcess, "goqueryd", now)
	v, e := s.SearchPlacements("")
	assert.FatalOnError(t, e)
	assert.Count(t, 2, v)
}

func TestSearchPlacementsRefusesAnAbsentName(t *testing.T) {
	s := store_tester.NewStore(t)
	store_tester.Save(t, s, constant.SourceKubernetes, "gitlab", time.Now())
	v, e := s.SearchPlacements("nowhere")
	assert.FatalOnError(t, e)
	assert.Count(t, 0, v)
}
