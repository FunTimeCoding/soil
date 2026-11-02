package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/login_session"
	"github.com/funtimecoding/soil/pkg/tool/gogated/unit/base"
	"github.com/funtimecoding/soil/pkg/worker"
	"testing"
	"time"
)

func TestOnlyTheSuperUserIsAdministrator(t *testing.T) {
	b := base.New(t)
	row, e := b.Store.UserByMail(constant.FixtureMail)
	assert.FatalOnError(t, e)
	assert.True(t, administrator(t, b, row.Identifier))
	other, f := b.Store.CreateDirectoryUser(
		"unique-9",
		"other",
		"other@example.test",
		"Other",
	)
	assert.FatalOnError(t, f)
	assert.False(t, administrator(t, b, other.Identifier))
	assert.False(t, administrator(t, b, "missing"))
	assert.False(t, administrator(t, b, ""))
}

func TestAFailingStoreIsNotAnAnswerAboutTheAdministrator(t *testing.T) {
	b := base.New(t)
	relational.CloseMapper(b.Database)
	_, e := b.Service.IsAdministrator("anyone")
	assert.Error(t, e)
}

func TestCleanDropsOnlyExpiredLoginSessions(t *testing.T) {
	b := base.New(t)
	stale, e := b.Service.CreateLoginSession("stale")
	assert.FatalOnError(t, e)
	fresh, f := b.Service.CreateLoginSession("fresh")
	assert.FatalOnError(t, f)
	assert.FatalOnError(
		t,
		b.Database.Model(login_session.Stub()).
			Where("identifier = ?", stale).
			Update(
				"created_at",
				time.Now().Add(-2*constant.LoginSessionTimeToLive),
			).Error,
	)
	assert.Count(t, 0, b.Service.Clean())
	var rows []login_session.LoginSession
	assert.FatalOnError(t, b.Database.Find(&rows).Error)
	assert.Count(t, 1, rows)
	assert.String(t, fresh, rows[0].Identifier)
}

func TestACleanTickReportsEachFailure(t *testing.T) {
	b := base.New(t)
	relational.CloseMapper(b.Database)
	r := memory.New()
	owned := worker.NewEvery(
		b.Service.Clean,
		logger.New(t.Context()),
		r,
		time.Hour,
	)
	defer owned.Close()
	owned.Poll()
	assert.Count(t, 2, r.Events())
}
