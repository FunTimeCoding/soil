package preposition_name

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func TestPrepositionNameFlagsTable(t *testing.T) {
	results := analyze(t)
	testutil.AssertBlocked(t, results, 42)

	for _, text := range []string{
		`OwnedBy carries "by"`,
		`From carries "from"`,
		`To carries "to"`,
		`PackOf carries "of"`,
		`ItemByName carries "by"`,
		`PriceOf carries "of"`,
		`valueAt carries "at"`,
		`scaledTo carries "to"`,
		`nameFromItem carries "from"`,
		`openAt carries "at"`,
		`RecentOrdersFor carries "for"`,
		`loadOn carries "on"`,
		`ReservedUntil carries "until"`,
		`SaveFor carries "for"`,
		`releaseFromList carries "from"`,
		`OpenedAt carries "at"`,
		`ShippedAt carries "at"`,
	} {
		testutil.AssertBlockedContains(t, results, text)
	}
}

func TestPrepositionNameLeavesBoundaries(t *testing.T) {
	results := analyze(t)

	for _, name := range []string{
		"CreatedAt ",
		"SeenAt ",
		"OnSaleUntil ",
		"ForSale ",
		"AtCapacity ",
		"RestockAt ",
		"ShipsAt ",
		"ClosedAt ",
		"AddToList ",
		"HeightToWidth ",
		"FromName ",
		"NewFromName ",
		"PanicOnError ",
		"OrdersPerDay ",
		"WithTimeout ",
		"FatalOnZero ",
		"TestItemByNameForEmptyStore ",
		"ClientByName ",
	} {
		testutil.AssertNotBlockedContains(t, results, name)
	}
}

func TestPrepositionNameMomentAnswerKeepsOnlyAt(t *testing.T) {
	results := analyze(t)
	testutil.AssertBlockedContains(t, results, `ResaleFrom carries "from"`)
	testutil.AssertBlockedContains(t, results, `dayOf carries "of"`)
	testutil.AssertNotBlockedContains(t, results, "RestockAt ")
}

func TestPrepositionNameChecksConstantsAndVariables(t *testing.T) {
	results := analyze(t)

	for _, text := range []string{
		`DaysOfYear carries "of"`,
		`AtDoorCondition carries "at"`,
		`PriceByLabel carries "by"`,
		`aisleOfDay carries "of"`,
	} {
		testutil.AssertBlockedContains(t, results, text)
	}

	for _, name := range []string{
		"ShelvedAtColumn ",
		"ShelvedAt ",
		"StockedAtIndex ",
	} {
		testutil.AssertNotBlockedContains(t, results, name)
	}
}

func TestPrepositionNameChecksParametersAndLocals(t *testing.T) {
	results := analyze(t)

	for _, text := range []string{
		`byName carries "by"`,
		`sinceOpen carries "since"`,
		`labelsByCrate carries "by"`,
		`labelByName carries "by"`,
	} {
		testutil.AssertBlockedContains(t, results, text)
	}

	for _, name := range []string{
		"from carries",
		"to carries",
		"since carries",
		"seenAt ",
	} {
		testutil.AssertNotBlockedContains(t, results, name)
	}
}

func TestPrepositionNameFlagsAnsweredWords(t *testing.T) {
	results := analyze(t)

	for _, text := range []string{
		`VersionsSince carries "since"`,
		`ItemsInAisle carries "in"`,
		`walkThrough carries "through"`,
		`RestockAfter carries "after"`,
		`SoldSince carries "since"`,
		`ExpiresBefore carries "before"`,
		`OrdersWithoutReceipt carries "without"`,
		`GapBetween carries "between"`,
		`CheapestAmong carries "among"`,
		`TurnToward carries "toward"`,
		`LabelsInside carries "inside"`,
		`LabelsOutside carries "outside"`,
		`CrateWithLabel carries "with"`,
	} {
		testutil.AssertBlockedContains(t, results, text)
	}

	for _, name := range []string{"WithNote ", "NewWithLabel ", "LookUp "} {
		testutil.AssertNotBlockedContains(t, results, name)
	}
}

func TestPrepositionNameSparesInOpeningQuestion(t *testing.T) {
	results := analyze(t)
	testutil.AssertNotBlockedContains(t, results, "InStock ")
	testutil.AssertBlockedContains(t, results, `InCart carries "in"`)
	testutil.AssertBlockedContains(t, results, `StockedInAisle carries "in"`)
}
