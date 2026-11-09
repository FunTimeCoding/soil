package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/client"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/unit/base"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func TestMockAnswersEveryCallAsDaemonDoes(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	seedServed(s.MockClient)
	c := client.New(
		locator.New("localhost").Port(s.ContextServer.Port).Insecure(),
		generativeConstant.ModelContextTestToken,
	)
	defer c.Close()
	m := mock_client.New()
	defer m.Close()
	seedMock(m)
	served := mediaAnswers(t.Context(), c)
	assert.Count(t, 19, served)
	assert.Strings(t, served, mediaAnswers(t.Context(), m))
}

func TestMockRecordsEveryWriteItAccepts(t *testing.T) {
	m := mock_client.New()
	defer m.Close()
	seedMock(m)
	mediaAnswers(t.Context(), m)
	assert.Count(t, 1, m.Played["session-1"])
	assert.Count(t, 1, m.Commanded["session-1"])
	assert.Integers(t, []int{40}, m.Volumes["session-1"])
}

func TestMockFailsOnlyMethodItIsToldTo(t *testing.T) {
	m := mock_client.New()
	defer m.Close()
	seedMock(m)
	m.Fail(m.GetItem, errors.New("item refused"))
	_, e := m.ListSessions(t.Context())
	assert.FatalOnError(t, e)
	_, f := m.GetItem(t.Context(), "movie-1")
	assert.String(t, "item refused", f.Error())
}
