package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/browser/face"
	"github.com/funtimecoding/soil/pkg/browser/ladybird"
	"github.com/funtimecoding/soil/pkg/browser/ladybird/mock_client"
	"github.com/funtimecoding/soil/pkg/browser/types/ladybird_option"
	"github.com/funtimecoding/soil/pkg/browser/unit/ladybird_tester"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"testing"
)

func TestLadybirdClientAndMockMatchFace(t *testing.T) {
	clients := []face.LadybirdClient{
		ladybird.New(ladybird_option.New()),
		mock_client.New(t),
	}
	assert.Count(t, 2, clients)
}

func TestLadybirdTreeDecodesContexts(t *testing.T) {
	l := ladybird_tester.New(t)
	l.Server.Answer(
		constant.TreeMethod,
		`{"contexts":[{"context":"c1","url":"about:blank","children":[]}]}`,
	)
	contexts, e := l.Client.Tree()
	assert.FatalOnError(t, e)
	assert.Count(t, 1, contexts)
	assert.String(t, "c1", contexts[0].Identifier)
	assert.String(t, "about:blank", contexts[0].Locator)
	assert.String(t, "browsingContext.getTree", l.Server.Last().Method)
}

func TestLadybirdMissingFrameNotFound(t *testing.T) {
	l := ladybird_tester.New(t)
	l.Server.Refuse(
		constant.NavigateMethod,
		constant.NoSuchFrame,
		"Frame not found",
	)
	e := l.Client.Navigate("c1", "https://example.com/")
	assert.True(t, not_found.Is(e))
	assert.String(
		t,
		"browsingContext.navigate: no such frame: Frame not found",
		e.Error(),
	)
}

func TestLadybirdInvalidArgumentValidation(t *testing.T) {
	l := ladybird_tester.New(t)
	l.Server.Refuse(
		constant.NavigateMethod,
		constant.InvalidArgument,
		"bad url",
	)
	assert.True(t, validation.Is(l.Client.Navigate("c1", "nowhere")))
}

func TestLadybirdUnknownCommandUnexpected(t *testing.T) {
	_, e := ladybird_tester.New(t).Client.Screenshot("c1")
	assert.True(t, unexpected.Is(e))
	assert.False(t, not_found.Is(e))
}

func TestLadybirdScriptExceptionValidation(t *testing.T) {
	l := ladybird_tester.New(t)
	l.Server.Answer(
		constant.EvaluateMethod,
		`{"type":"exception","exceptionDetails":{"text":"ReferenceError: x"}}`,
	)
	_, e := l.Client.Evaluate("c1", "x")
	assert.True(t, validation.Is(e))
	assert.String(t, "script exception: ReferenceError: x", e.Error())
}

func TestLadybirdEvaluateReturnsResult(t *testing.T) {
	l := ladybird_tester.New(t)
	l.Server.Answer(
		constant.EvaluateMethod,
		`{"type":"success","result":{"type":"number","value":3}}`,
	)
	result, e := l.Client.Evaluate("c1", "1+2")
	assert.FatalOnError(t, e)
	assert.String(t, `{"type":"number","value":3}`, string(result))
}

func TestLadybirdScreenshotDecodesImage(t *testing.T) {
	l := ladybird_tester.New(t)
	l.Server.Answer(constant.CaptureMethod, `{"data":"cG5n"}`)
	image, e := l.Client.Screenshot("c1")
	assert.FatalOnError(t, e)
	assert.String(t, "png", string(image))
}

func TestLadybirdEventReachesCallback(t *testing.T) {
	l := ladybird_tester.New(t)
	l.Server.Answer(constant.SubscribeMethod, `{}`)
	assert.FatalOnError(t, l.Client.Subscribe("log.entryAdded"))
	l.Server.Emit("log.entryAdded", `{"level":"info","text":"hello"}`)
	r := l.Event()
	assert.String(t, "log.entryAdded", r.Method)
	assert.String(t, `{"level":"info","text":"hello"}`, string(r.Parameters))
}

func TestLadybirdCloseDeletesSession(t *testing.T) {
	l := ladybird_tester.New(t)
	l.Client.Close()
	assert.True(t, l.Server.Deleted())
}

func TestLadybirdCloseSurvivesStoppedDriver(t *testing.T) {
	l := ladybird_tester.New(t)
	l.Server.Stop()
	l.Client.Close()
	assert.False(t, l.Server.Deleted())
}

func TestLadybirdTreeAfterCloseUnavailable(t *testing.T) {
	l := ladybird_tester.New(t)
	l.Client.Close()
	_, e := l.Client.Tree()
	assert.True(t, unavailable.Is(e))
}
