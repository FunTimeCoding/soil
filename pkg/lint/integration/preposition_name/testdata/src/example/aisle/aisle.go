package aisle

import "time"

// golint:fixture stray_constant
const (
	ShelvedAtColumn = "shelved_at"
	ShelvedAt       = "shelved"
	AtDoorCondition = "door"
	PriceByLabel    = "label"
)

var (
	StockedAtIndex = "stocked_at_index"
	aisleOfDay     = 3
)

type Crate struct {
	Labels []string
	Open   bool
}

func ItemsInAisle(aisle int) int {
	return aisle + aisleOfDay
}

func (c *Crate) InStock(label string) bool {
	return len(c.Labels) > 0 && label != ""
}

func (c *Crate) InCart(label string) int {
	return len(c.Labels) + len(label)
}

func (c *Crate) StockedInAisle(label string) bool {
	return len(c.Labels) > len(label)
}

func walkThrough(c *Crate) int {
	return len(c.Labels)
}

func RestockAfter(c *Crate) int {
	return len(c.Labels)
}

func SoldSince(c *Crate) int {
	return len(c.Labels)
}

func ExpiresBefore(c *Crate) int {
	return len(c.Labels)
}

func OrdersWithoutReceipt(c *Crate) int {
	return len(c.Labels)
}

func GapBetween(
	a int,
	b int,
) int {
	return a - b
}

func CheapestAmong(c *Crate) int {
	return len(c.Labels)
}

func TurnToward(c *Crate) int {
	return len(c.Labels)
}

func LabelsInside(c *Crate) int {
	return len(c.Labels)
}

func LabelsOutside(c *Crate) int {
	return len(c.Labels)
}

func CrateWithLabel(c *Crate) int {
	return len(c.Labels)
}

func WithNote(note string) *Crate {
	return &Crate{Labels: []string{note}}
}

func NewWithLabel(label string) *Crate {
	return &Crate{Labels: []string{label}}
}

func LookUp(label string) int {
	return len(label)
}

func dayOf(t time.Time) time.Time {
	return t.Truncate(24 * time.Hour)
}

func count(
	c *Crate,
	from int,
	to int,
) int {
	sinceOpen := to - from
	seenAt := time.Now()
	labelsByCrate := map[*Crate]int{c: len(c.Labels)}

	return sinceOpen + seenAt.Second() + labelsByCrate[c]
}

func pick(
	c *Crate,
	labelByName string,
) string {
	return c.Labels[0] + labelByName
}
