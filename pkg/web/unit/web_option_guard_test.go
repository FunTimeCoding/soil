package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/web/authorization/authorization_tester"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/option"
	"github.com/funtimecoding/soil/pkg/web/unit/web_tester"
	"net/http"
	"testing"
)

func TestGuardMountsSignOnOnlyWhenOptionCarriesIt(t *testing.T) {
	signed := option.New("", "token")
	signed.Authorization = authorization_tester.Client
	m := http.NewServeMux()
	g, release := signed.Guard(m, memory.New())
	t.Cleanup(release)
	assert.True(t, g.HasSession())
	assert.Integer(t, http.StatusFound, served(m, constant.SignInPath).Code)
	plain := http.NewServeMux()
	p, done := option.New("", "token").Guard(plain, memory.New())
	t.Cleanup(done)
	assert.False(t, p.HasSession())
	assert.Integer(
		t,
		http.StatusNotFound,
		served(plain, constant.SignInPath).Code,
	)
	assert.Panic(t, func() { p.Session("GET /target", web_tester.Serve) })
}
