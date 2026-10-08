package record

type LegacyEvent struct {
	Identifier uint   `gorm:"column:identifier"`
	Kind       string `gorm:"column:kind"`
	Scope      string `gorm:"column:scope"`
	Body       string `gorm:"column:body"`
}
