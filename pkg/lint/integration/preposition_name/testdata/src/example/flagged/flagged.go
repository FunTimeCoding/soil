package flagged

import "time"

type Item struct {
	Name    string
	OwnedBy string
	From    int
	To      int
}

type PackOf struct {
	Left string
}

type Owner interface {
	ItemByName(name string) *Item
}

func PriceOf(i *Item) int {
	return len(i.Name)
}

func valueAt(
	i *Item,
	second int,
) int {
	return len(i.Name) + second
}

func scaledTo(
	i *Item,
	size int,
) int {
	return len(i.Name) + size
}

func nameFromItem(i *Item) string {
	return i.Name
}

func openAt(
	i *Item,
	second int,
) bool {
	return len(i.Name) > second
}

func RecentOrdersFor(name string) []string {
	return []string{name}
}

func loadOn(i *Item) int {
	return len(i.Name)
}

func ReservedUntil(i *Item) int {
	return len(i.Name)
}

func SaveFor(i *Item) error {
	if i.Name == "" {
		return nil
	}

	return nil
}

func releaseFromList(
	i *Item,
	list []*Item,
) {
	list[0] = i
}

func OpenedAt(i *Item) (time.Time, int) {
	return time.Time{}, len(i.Name)
}

func ShippedAt(i *Item) (time.Time, bool, int) {
	return time.Time{}, true, len(i.Name)
}
