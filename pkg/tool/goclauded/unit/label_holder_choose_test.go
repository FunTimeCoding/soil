package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/label_holder_tester"
	"testing"
	"time"
)

func TestHolderKeepsTheLabel(t *testing.T) {
	plan := label_holder_tester.Choose(
		label_holder_tester.Active("a", ""),
		label_holder_tester.Active("b", "watch"),
	)
	assert.String(t, "b", plan.Holder)
	assert.String(t, "", plan.Adopt)
	assert.Count(t, 0, plan.Unset)
}

func TestActiveHolderKeepsAgainstAResumedOne(t *testing.T) {
	plan := label_holder_tester.Choose(
		label_holder_tester.Idle("resumed", "watch"),
		label_holder_tester.Active("holder", "watch"),
	)
	assert.String(t, "holder", plan.Holder)
	assert.Count(t, 1, plan.Unset)
	assert.String(t, "resumed", plan.Unset[0].Identifier)
}

func TestIdleHolderIsReplaced(t *testing.T) {
	plan := label_holder_tester.Choose(
		label_holder_tester.Active("fresh", ""),
		label_holder_tester.Idle("gone", "watch"),
	)
	assert.String(t, "fresh", plan.Adopt)
	assert.Count(t, 1, plan.Unset)
	assert.String(t, "gone", plan.Unset[0].Identifier)
}

func TestDuplicatesTrimToTheNewest(t *testing.T) {
	plan := label_holder_tester.Choose(
		label_holder_tester.Active("a", "watch"),
		label_holder_tester.Active("b", "watch"),
	)
	assert.String(t, "a", plan.Holder)
	assert.Count(t, 1, plan.Unset)
	assert.String(t, "b", plan.Unset[0].Identifier)
}

func TestAdoptionSkipsSessionsHoldingAnotherValue(t *testing.T) {
	plan := label_holder_tester.Choose(
		label_holder_tester.Active("a", "bravo"),
		label_holder_tester.Active("b", ""),
		label_holder_tester.Active("c", ""),
	)
	assert.String(t, "b", plan.Adopt)
	assert.Count(t, 0, plan.Unset)
}

func TestIdleSessionsAreNeverAdopted(t *testing.T) {
	plan := label_holder_tester.Choose(label_holder_tester.Idle("a", ""))
	assert.String(t, "", plan.Holder)
}

func TestValueMatchesCaseInsensitively(t *testing.T) {
	plan := label_holder_tester.Choose(label_holder_tester.Active("a", "Watch"))
	assert.String(t, "a", plan.Holder)
	assert.String(t, "", plan.Adopt)
}

func TestEveryoneDeclinedHoldsNobody(t *testing.T) {
	plan := label_holder_tester.Choose(
		label_holder_tester.Active("a", "none"),
		label_holder_tester.Active("b", "bravo"),
	)
	assert.String(t, "", plan.Holder)
	assert.String(t, "", plan.Adopt)
}

func TestFreshSessionsWaitForTheirTurns(t *testing.T) {
	plan := label_holder_tester.Choose(
		label_holder_tester.Seen("young", "", time.Minute, 2),
		label_holder_tester.Active("settled", ""),
	)
	assert.String(t, "settled", plan.Adopt)
}

func TestEmptyRosterHoldsNobody(t *testing.T) {
	plan := label_holder_tester.Choose()
	assert.String(t, "", plan.Holder)
	assert.Count(t, 0, plan.Unset)
}
