package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/band"
	"github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"io"
	"net/http"
	"testing"
)

func TestListenerPutBindsTheEnvelopePrefix(t *testing.T) {
	s, last := upstream_tester.New(
		t,
		http.StatusOK,
		constant.FixtureRedirection,
	)
	c := band.New(s.URL, "alfa", "bravo")
	defer c.Close()
	assert.FatalOnError(t, c.EnableListener())
	b, e := io.ReadAll(last().Body)
	assert.FatalOnError(t, e)
	put := string(b)
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
