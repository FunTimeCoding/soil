package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/gitlab/mock_client"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/service"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/filter"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/server_tester"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/store_tester"
	"testing"
)

func TestTheServiceClosesTheDatabaseItOpened(t *testing.T) {
	v := service.New(
		store_tester.NewOption(t, mock_client.New()),
		logger.New(t.Context()),
		memory.New(),
	)
	_, e := v.Certificates(filter.New())
	assert.FatalOnError(t, e)
	v.Close()
	v.Close()
	_, f := v.Certificates(filter.New())
	assert.FatalNotNil(t, f)
	assert.StringContains(t, "database is closed", f.Error())
}

func TestRevokingTwiceFindsNothingTheSecondTime(t *testing.T) {
	v := store_tester.NewChain(t)
	r, _, e := v.IssueCertificate(
		server_tester.LeafBody(
			constant.FixtureCommonName,
			[]string{constant.FixtureHost},
		),
	)
	assert.FatalOnError(t, e)
	first, f := v.Revoke(r.Serial)
	assert.FatalOnError(t, f)
	assert.NotNil(t, first.Revoked)
	_, g := v.Revoke(r.Serial)
	assert.True(t, not_found.Is(g))
}
