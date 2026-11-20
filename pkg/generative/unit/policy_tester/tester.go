package policy_tester

import "testing"

type Tester struct {
	Root       string
	Scratchpad string
	t          *testing.T
}
