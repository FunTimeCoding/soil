package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/client"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/client/mock_client"
	generated "github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/unit/base"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func TestMockAnswersDirectoryAsDaemonDoes(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	c := client.New(
		locator.New("localhost").Port(s.Port).Insecure(),
		generativeConstant.ModelContextTestToken,
	)
	defer c.Close()
	m := mock_client.New()
	defer m.Close()
	served := directoryAnswers(t.Context(), c)
	assert.Count(t, 25, served)
	assert.Strings(t, served, directoryAnswers(t.Context(), m))
}

func TestMockKeepsPasswordItWasGiven(t *testing.T) {
	m := mock_client.New()
	defer m.Close()
	directoryAnswers(t.Context(), m)
	assert.Count(t, 0, m.Passwords)
	_, e := m.PostUser(
		t.Context(),
		generated.PostUserJSONRequestBody{
			Account:  "golf",
			Password: new("hotel"),
		},
	)
	assert.FatalOnError(t, e)
	assert.String(t, "hotel", m.Passwords["golf"])
}

func TestMockFailsOnlyMethodItIsToldTo(t *testing.T) {
	m := mock_client.New()
	defer m.Close()
	m.Fail(m.GetUserAccount, errors.New("account refused"))
	_, e := m.GetUser(t.Context())
	assert.FatalOnError(t, e)
	_, f := m.GetUserAccount(t.Context(), "alfa")
	assert.String(t, "account refused", f.Error())
}
