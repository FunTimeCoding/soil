package outside

import (
	"example/flagged"
	"strings"
	"time"
)

// golint:fixture stray_constant
const DaysOfYear = 365

type Item struct {
	Name      string
	CreatedAt time.Time
	SeenAt      string
	OnSaleUntil *time.Time
	ForSale     bool
}

type Store struct {
	items []*flagged.Item
}

func (s *Store) ItemByName(name string) *flagged.Item {
	for _, i := range s.items {
		if i.Name == name {
			return i
		}
	}

	return nil
}

func (i *Item) AtCapacity() bool {
	return i.ForSale
}

func AddToList(
	i *Item,
	list []*Item,
) []*Item {
	return append(list, i)
}

func HeightToWidth(i *Item) float64 {
	return float64(len(i.Name))
}

func FromName(name string) *Item {
	return &Item{Name: name}
}

func NewFromName(name string) *Item {
	return &Item{Name: name}
}

func PanicOnError(e error) {
	if e != nil {
		panic(e)
	}
}

func (i *Item) RestockAt() time.Time {
	return i.CreatedAt
}

func ShipsAt(i *Item) (time.Time, bool) {
	return i.CreatedAt, i.ForSale
}

func ResaleFrom(i *Item) (*time.Time, error) {
	return i.OnSaleUntil, nil
}

func ClosedAt(i *Item) (time.Time, bool, error) {
	return i.CreatedAt, i.ForSale, nil
}

func VersionsSince(
	name string,
	since int,
) []string {
	return []string{name, strings.Repeat("x", since)}
}

func OrdersPerDay(i *Item) int {
	return len(i.Name)
}

func WithTimeout(seconds int) int {
	return seconds
}

func FatalOnZero(v int) int {
	return v
}

func describe(i *Item) string {
	byName := i.Name
	from := strings.ToUpper(byName)

	return from
}
