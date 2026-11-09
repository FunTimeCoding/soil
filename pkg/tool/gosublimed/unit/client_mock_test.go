package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/client"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/unit/base"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func TestMockAnswersEveryViewAsDaemonDoes(t *testing.T) {
	s := base.New(t)
	c := client.New(
		locator.New("localhost").Port(s.Port).Insecure(),
		generativeConstant.ModelContextTestToken,
	)
	defer c.Close()
	m := mock_client.New()
	defer m.Close()
	served := viewAnswers(t.Context(), c)
	assert.Count(t, 19, served)
	assert.Strings(t, served, viewAnswers(t.Context(), m))
}

func TestMockFailsOnlyMethodItIsToldTo(t *testing.T) {
	m := mock_client.New()
	defer m.Close()
	m.Fail(m.GetView, errors.New("view refused"))
	_, e := m.GetViews(t.Context())
	assert.FatalOnError(t, e)
	_, f := m.GetView(t.Context(), 1)
	assert.String(t, "view refused", f.Error())
}
