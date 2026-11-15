package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/band"
	"github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"testing"
)

func TestListenerPutBindsEnvelopePrefix(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer(constant.FixtureRedirection))
	c := band.New(s.URL, "alfa", "bravo")
	defer c.Close()
	assert.FatalOnError(t, c.EnableListener())
	put := s.LastBody()
	assert.StringContains(
		t,
		"http://schemas.xmlsoap.org/ws/2004/09/transfer/Put",
		put,
	)
	assert.StringContains(
		t,
		`<s:Body><g:AMT_RedirectionService xmlns:g="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RedirectionService"><g:CreationClassName>AMT_RedirectionService</g:CreationClassName><g:ElementName>Intel(r) AMT Redirection Service</g:ElementName><g:EnabledState>32771</g:EnabledState><g:ListenerEnabled>true</g:ListenerEnabled><g:Name>Intel(r) AMT Redirection Service</g:Name><g:SystemCreationClassName>CIM_ComputerSystem</g:SystemCreationClassName><g:SystemName>Intel(r) AMT</g:SystemName></g:AMT_RedirectionService></s:Body>`,
		put,
	)
}
