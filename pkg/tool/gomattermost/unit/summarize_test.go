package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	ollamaMock "github.com/funtimecoding/soil/pkg/generative/ollama/mock_client"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/gomattermost"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/client"
	"strings"
	"testing"
	"time"
)

func TestEachThreadCarriesTheSummaryOfItsOwnLog(t *testing.T) {
	c := mock_client.New()
	defer c.Close()
	c.Unresolved["juliett"] = []client.Thread{
		{Root: client.Post{Username: "lima", Message: "alfa is down"}},
		{Root: client.Post{Username: "mike", Message: "bravo is slow"}},
	}
	o := ollamaMock.New()
	defer o.Close()
	o.Respond = func(prompt string) string {
		if strings.Contains(prompt, "alfa is down") {
			return "kilo"
		}

		return "papa"
	}
	result, e := gomattermost.Summarize(c, o, "juliett", time.Now())
	assert.FatalOnError(t, e)
	assert.Count(t, 2, result)
	assert.String(t, "alfa is down", result[0].Thread.Root.Message)
	assert.String(t, "kilo", result[0].Summary)
	assert.String(t, "bravo is slow", result[1].Thread.Root.Message)
	assert.String(t, "papa", result[1].Summary)
	assert.Count(t, 2, o.Prompts)
	assert.StringContains(
		t,
		"Summarize this tech support chat thread.",
		o.Prompts[0],
	)
}

func TestNothingUnresolvedIsAnEmptyList(t *testing.T) {
	c := mock_client.New()
	defer c.Close()
	c.Unresolved["juliett"] = []client.Thread{}
	o := ollamaMock.New()
	defer o.Close()
	result, e := gomattermost.Summarize(c, o, "juliett", time.Now())
	assert.FatalOnError(t, e)
	assert.String(t, "[]", string(notation.Marshal(result)))
	assert.Count(t, 0, o.Prompts)
}

func TestAnUnavailableModelIsReturnedAsAnError(t *testing.T) {
	c := mock_client.New()
	defer c.Close()
	c.Unresolved["juliett"] = []client.Thread{
		{Root: client.Post{Username: "lima", Message: "alfa is down"}},
	}
	o := ollamaMock.New()
	defer o.Close()
	o.Failure = errors.New("foxtrot")
	result, e := gomattermost.Summarize(c, o, "juliett", time.Now())
	assert.FatalNotNil(t, e)
	assert.StringContains(t, "foxtrot", e.Error())
	assert.Nil(t, result)
	assert.Count(t, 1, o.Prompts)
}

func TestAFailedThreadListAsksTheModelNothing(t *testing.T) {
	c := mock_client.New()
	defer c.Close()
	c.Failure = errors.New("foxtrot")
	o := ollamaMock.New()
	defer o.Close()
	_, e := gomattermost.Summarize(c, o, "juliett", time.Now())
	assert.NotNil(t, e)
	assert.Count(t, 0, o.Prompts)
}
