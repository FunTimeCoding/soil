package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	generated "github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/server"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/unit/service_tester"
	"net"
	"os"
	"syscall"
	"testing"
)

func TestAnOutageIsAnsweredAsUnavailableWithoutACapture(t *testing.T) {
	c := service_tester.NewSource(t)
	c.Failure = &net.OpError{
		Op:  "dial",
		Net: "tcp",
		Err: os.NewSyscallError("connect", syscall.ECONNREFUSED),
	}
	m := memory.New()
	s := server.New(service_tester.Service(t, c, m), m)

	for range 2 {
		r, e := s.GetLists(
			context.Background(),
			generated.GetListsRequestObject{},
		)
		assert.FatalOnError(t, e)
		_, away := r.(generated.GetLists503JSONResponse)
		assert.True(t, away)
	}

	assert.Count(t, 0, m.Events())
}

func TestARefusedTokenIsAnsweredAsForbiddenWithoutACapture(t *testing.T) {
	c := service_tester.NewSource(t)
	c.Failure = forbidden.Format("linkace: %s", "Unauthenticated.")
	m := memory.New()
	r, e := server.New(service_tester.Service(t, c, m), m).GetTags(
		context.Background(),
		generated.GetTagsRequestObject{},
	)
	assert.FatalOnError(t, e)
	_, refused := r.(generated.GetTags403JSONResponse)
	assert.True(t, refused)
	assert.Count(t, 0, m.Events())
}

func TestARejectedLinkIsAnsweredAsBadRequestWithoutACapture(t *testing.T) {
	c := service_tester.NewSource(t)
	c.Failure = validation.New("linkace: %s", "The url field is required.")
	m := memory.New()
	r, e := server.New(service_tester.Service(t, c, m), m).AddLink(
		context.Background(),
		generated.AddLinkRequestObject{
			Body: &generated.AddLinkJSONRequestBody{Link: "not a url"},
		},
	)
	assert.FatalOnError(t, e)
	_, rejected := r.(generated.AddLink400JSONResponse)
	assert.True(t, rejected)
	assert.Count(t, 0, m.Events())
}

func TestATakenTagNameIsAnsweredAsConflictWithoutACapture(t *testing.T) {
	c := service_tester.NewSource(t)
	c.Failure = conflict.Format("linkace: %s", "The tag already exists.")
	m := memory.New()
	r, e := server.New(service_tester.Service(t, c, m), m).CreateTag(
		context.Background(),
		generated.CreateTagRequestObject{
			Body: &generated.CreateTagJSONRequestBody{Name: "deploy"},
		},
	)
	assert.FatalOnError(t, e)
	_, taken := r.(generated.CreateTag409JSONResponse)
	assert.True(t, taken)
	assert.Count(t, 0, m.Events())
}
