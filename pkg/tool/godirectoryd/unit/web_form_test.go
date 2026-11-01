package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/unit/base"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"io"
	"net/http"
	"net/url"
	"testing"
)

func TestAFailedFormKeepsItsAlertAndValuesAtTheForm(t *testing.T) {
	s := base.New(t)
	_, e := s.Service.CreateGroup("alfa")
	assert.FatalOnError(t, e)
	q := web.NewPost(
		s.Address(constant.CreateGroupPath),
		url.Values{constant.NameField: {"alfa"}}.Encode(),
	)
	q.Header.Set(webConstant.ContentType, webConstant.FormEncoded)
	web.Bearer(q, generativeConstant.ModelContextTestToken)
	r, f := s.Halting().Do(q)
	assert.FatalOnError(t, f)
	errors.PanicClose(r.Body)
	assert.Integer(t, http.StatusSeeOther, r.StatusCode)
	page := web.NewGet(s.Address(r.Header.Get("Location")))
	web.Bearer(page, generativeConstant.ModelContextTestToken)
	g, h := s.Halting().Do(page)
	assert.FatalOnError(t, h)
	defer errors.PanicClose(g.Body)
	b, i := io.ReadAll(g.Body)
	assert.FatalOnError(t, i)
	body := string(b)
	assert.StringContains(
		t,
		`<form class="entry" id="create-group" method="post" action="/group/create"><h5>Create group</h5><p role="alert" class="form-alert">`,
		body,
	)
	assert.StringContains(t, `name="name" value="alfa"`, body)
	assert.StringContains(
		t,
		`<form class="entry" id="membership" method="post" action="/group/member/add"><h5>Membership</h5><label>`,
		body,
	)
}
