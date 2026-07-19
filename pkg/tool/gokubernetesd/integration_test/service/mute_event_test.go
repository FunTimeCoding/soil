package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/integration_test/service_tester"
	"testing"
)

func TestMuteEvent(t *testing.T) {
	s := service_tester.New(t)
	e := s.Service.MuteEvent("DNSConfigForming", "", "")
	assert.Nil(t, e)
	rules := s.MuteRules()
	assert.Count(t, 1, rules)
	assert.String(t, "DNSConfigForming", rules[0].Reason)
}

func TestMuteEventWithCluster(t *testing.T) {
	s := service_tester.New(t)
	s.Mute("BackOff", "", "test")
	rules := s.MuteRules()
	assert.Count(t, 1, rules)
	assert.String(t, "test", rules[0].Cluster)
}

func TestUnmuteEvent(t *testing.T) {
	s := service_tester.New(t)
	s.Mute("DNSConfigForming", "", "")
	rules := s.MuteRules()
	e := s.Service.UnmuteEvent(rules[0].Identifier)
	assert.Nil(t, e)
	assert.Count(t, 0, s.MuteRules())
}

func TestIsMuted(t *testing.T) {
	s := service_tester.New(t)
	s.Mute("DNSConfigForming", "", "")
	assert.True(t, s.IsMuted("DNSConfigForming", "some message", "test"))
}

func TestIsMutedCaseInsensitive(t *testing.T) {
	s := service_tester.New(t)
	s.Mute("DNSConfigForming", "nameserver", "")
	assert.True(
		t,
		s.IsMuted("DNSConfigForming", "Nameserver limits exceeded", "test"),
	)
	assert.False(t, s.IsMuted("DNSConfigForming", "something else", "test"))
}

func TestIsMutedClusterScoped(t *testing.T) {
	s := service_tester.New(t)
	s.Mute("BackOff", "", "production")
	assert.True(t, s.IsMuted("BackOff", "", "production"))
	assert.False(t, s.IsMuted("BackOff", "", "staging"))
}
