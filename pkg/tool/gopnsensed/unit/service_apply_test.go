package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/opnsense/request"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/mock_source"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/unit/counting_source"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/unit/rejecting_host"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/unit/service_tester"
	"testing"
)

func TestAnAppliedWriteReconfiguresDnsmasq(t *testing.T) {
	c := counting_source.New(mock_source.New())
	s := service_tester.New(t, c)
	_, e := s.AddHost(request.New(), true)
	assert.FatalOnError(t, e)
	assert.FatalOnError(t, s.SetHost("alfa", request.New(), true))
	assert.FatalOnError(t, s.DeleteHost("alfa", true))
	assert.Integer(t, 3, c.Reconfigured())
}

func TestADeferredWriteLeavesDnsmasqAlone(t *testing.T) {
	c := counting_source.New(mock_source.New())
	s := service_tester.New(t, c)
	_, e := s.AddHost(request.New(), false)
	assert.FatalOnError(t, e)
	assert.FatalOnError(t, s.SetHost("alfa", request.New(), false))
	assert.FatalOnError(t, s.DeleteHost("alfa", false))
	assert.Integer(t, 0, c.Reconfigured())
}

func TestARejectedWriteLeavesDnsmasqAlone(t *testing.T) {
	c := counting_source.New(rejecting_host.New())
	_, e := service_tester.New(t, c).AddHost(request.New(), true)
	assert.Error(t, e)
	assert.Integer(t, 0, c.Reconfigured())
}
