package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/service_tester"
	"testing"
)

func TestProfileStaysWithinByteBudget(t *testing.T) {
	o := service_tester.New(t)
	createTagged(t, o, "standing rule", constant.AlwaysTag)
	createBulky(t, o)
	result, d, e := o.Service.Profile("", "", true)
	assert.FatalOnError(t, e)
	assert.Greater(t, 0, d.IndexTrimmed)
	assert.Integer(t, len(result.Text), d.TotalBytes)
	assert.LessEqual(t, 50000, len(result.Text))
	assert.Greater(t, 49000, len(result.Text))
}

func TestProfileChargesNoticeToByteBudget(t *testing.T) {
	o := service_tester.New(t)
	createTagged(t, o, "standing rule", constant.AlwaysTag)
	createBulky(t, o)
	o.Indexer.Failure = unavailable.Format("goqueryd down")
	result, d, e := o.Service.Profile("retry", "", true)
	assert.FatalOnError(t, e)
	assert.String(
		t,
		"goqueryd is away - topic-matched memories and completions left out",
		result.Notice,
	)
	assert.Integer(t, len(result.Text), d.TotalBytes)
	assert.LessEqual(t, 50000, len(result.Text))
}
