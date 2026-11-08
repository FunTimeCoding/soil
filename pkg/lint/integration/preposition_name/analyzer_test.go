package preposition_name

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func TestPrepositionNameFlagsTheTable(t *testing.T) {
	results := analyze(t)
	testutil.AssertBlocked(t, results, 17)

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

func TestPrepositionNameLeavesTheBoundaries(t *testing.T) {
	results := analyze(t)

	for _, name := range []string{
		"DaysOfYear ",
		"CreatedAt ",
		"SeenAt ",
		"OnSaleUntil ",
		"ForSale ",
		"AtCapacity ",
		"RestockAt ",
		"ShipsAt ",
		"ResaleFrom ",
		"ClosedAt ",
		"AddToList ",
		"HeightToWidth ",
		"FromName ",
		"NewFromName ",
		"PanicOnError ",
		"VersionsSince ",
		"OrdersPerDay ",
		"WithTimeout ",
		"FatalOnZero ",
		"byName ",
		"TestItemByNameForEmptyStore ",
		"ClientByName ",
	} {
		testutil.AssertNotBlockedContains(t, results, name)
	}
}
