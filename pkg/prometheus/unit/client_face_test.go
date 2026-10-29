package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/prometheus"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/face"
	"github.com/funtimecoding/soil/pkg/prometheus/mock_client"
	"github.com/funtimecoding/soil/pkg/prometheus/query_result"
	"testing"
	"time"
)

func TestTheClientAndItsMockShareTheFace(t *testing.T) {
	c := prometheus.New("127.0.0.1", 0, false, "", "")
	t.Cleanup(c.Close)
	m := mock_client.New()
	t.Cleanup(m.Close)
	faces := []face.Client{c, m}
	assert.Count(t, 2, faces)
}

func TestTheMockAnswersWhatWasSeeded(t *testing.T) {
	m := mock_client.New()
	t.Cleanup(m.Close)
	r := query_result.New(nil, nil, nil)
	m.SeedResult(constant.Up, r)
	m.SeedFloat("free", 0.25)
	result, e := m.Query(constant.Up, time.Now())
	assert.FatalOnError(t, e)
	assert.True(t, result == r)
	value, f := m.QueryFloat("free", time.Now())
	assert.FatalOnError(t, f)
	assert.Float(t, 0.25, value)
}

func TestTheMockFailsEveryCallOnFailure(t *testing.T) {
	m := mock_client.New()
	t.Cleanup(m.Close)
	m.SeedFloat("free", 0.25)
	m.Failure = unavailable.Format("prometheus maintenance")
	_, e := m.QueryFloat("free", time.Now())
	assert.True(t, unavailable.Is(e))
	_, f := m.Query(constant.Up, time.Now())
	assert.True(t, unavailable.Is(f))
}
