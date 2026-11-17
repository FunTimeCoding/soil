package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/service_tester"
	"testing"
)

func TestProfileStaysWithinByteLimit(t *testing.T) {
	o := service_tester.New(t)
	o.CreateTagged("standing rule", constant.AlwaysTag)
	o.CreateBulky()
	result, d, e := o.Service().Profile("", "", true)
	assert.FatalOnError(t, e)
	assert.Greater(t, 0, d.IndexTrimmed)
	assert.Integer(t, len(result.Text), d.TotalBytes)
	assert.LessEqual(t, 50000, len(result.Text))
	assert.Greater(t, 49000, len(result.Text))
}

func TestProfileCountsNoticeAgainstByteLimit(t *testing.T) {
	o := service_tester.New(t)
	o.CreateTagged("standing rule", constant.AlwaysTag)
	o.CreateBulky()
	o.Indexer.Failure = unavailable.Format("goqueryd down")
	result, d, e := o.Service().Profile("retry", "", true)
	assert.FatalOnError(t, e)
	assert.String(
		t,
		"goqueryd is away - topic-matched memories and completions left out",
		result.Notice,
	)
	assert.Integer(t, len(result.Text), d.TotalBytes)
	assert.LessEqual(t, 50000, len(result.Text))
}
