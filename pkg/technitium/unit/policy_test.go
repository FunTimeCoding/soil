package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/technitium/policy"
	"testing"
)

func TestPolicyEncodeJoinsRowsWithCommaSeparatedTypes(t *testing.T) {
	assert.String(
		t,
		"alfa|*.alfa.example|A,TXT|bravo|bravo.example|AAAA",
		policy.Encode(
			[]*policy.Policy{
				policy.New("alfa", "*.alfa.example", []string{"A", "TXT"}),
				policy.New("bravo", "bravo.example", []string{"AAAA"}),
			},
		),
	)
}

func TestPolicyEncodeOfNothingClearsTable(t *testing.T) {
	assert.String(t, "false", policy.Encode(nil))
}

func TestPolicyUpsertReplacesSameKeyAndDomain(t *testing.T) {
	v := policy.Upsert(
		[]*policy.Policy{policy.New("alfa", "*.alfa.example", []string{"A"})},
		policy.New("alfa", "*.alfa.example", []string{"A", "TXT"}),
	)
	assert.Count(t, 1, v)
	assert.Strings(t, []string{"A", "TXT"}, v[0].Types)
}

func TestPolicyUpsertAppendsAnotherDomain(t *testing.T) {
	assert.Count(
		t,
		2,
		policy.Upsert(
			[]*policy.Policy{
				policy.New("alfa", "*.alfa.example", []string{"A"}),
			},
			policy.New("alfa", "alfa.example", []string{"A"}),
		),
	)
}
