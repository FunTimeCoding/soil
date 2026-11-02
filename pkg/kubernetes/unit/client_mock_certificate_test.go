package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/mock_client"
	"k8s.io/api/networking/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestTheMockIssuesATriggeredCertificateOnlyWhenToldTo(t *testing.T) {
	c := mock_client.New()
	c.AddCertificate("alfa", "bravo", "charlie")
	assert.FatalOnError(t, c.TriggerCertificateIssuance("alfa", "bravo"))
	pending, e := c.Certificates()
	assert.FatalOnError(t, e)
	assert.True(t, pending[0].Issuing)
	assert.False(t, pending[0].Ready)
	c.IssueOnTrigger()
	assert.FatalOnError(t, c.TriggerCertificateIssuance("alfa", "bravo"))
	issued, f := c.Certificates()
	assert.FatalOnError(t, f)
	assert.True(t, issued[0].Ready)
	assert.False(t, issued[0].Issuing)
	assert.Any(t, []string{"alfa/bravo", "alfa/bravo"}, c.Triggered())
	assert.True(t, not_found.Is(c.TriggerCertificateIssuance("alfa", "delta")))
}

func TestTheMockStoresTheNetworkPolicyItIsSent(t *testing.T) {
	c := mock_client.New()
	c.AddNetworkPolicy(
		&v1.NetworkPolicy{
			ObjectMeta: meta.ObjectMeta{Namespace: "alfa", Name: "bravo"},
		},
	)
	p, e := c.NetworkPolicy("alfa", "bravo")
	assert.FatalOnError(t, e)
	p.Spec.Ingress = append(p.Spec.Ingress, v1.NetworkPolicyIngressRule{})
	assert.FatalOnError(t, c.UpdateNetworkPolicy(p))
	q, f := c.NetworkPolicy("alfa", "bravo")
	assert.FatalOnError(t, f)
	assert.Count(t, 1, q.Spec.Ingress)
	_, g := c.NetworkPolicy("alfa", "charlie")
	assert.True(t, not_found.Is(g))
}

func TestTheMockFailsOneMethodAlone(t *testing.T) {
	c := mock_client.New()
	c.AddCertificate("alfa", "bravo", "charlie")
	c.Fail(c.TriggerCertificateIssuance, unexpected.Format("refused"))
	_, e := c.Certificates()
	assert.FatalOnError(t, e)
	assert.String(
		t,
		"refused",
		c.TriggerCertificateIssuance("alfa", "bravo").Error(),
	)
}
