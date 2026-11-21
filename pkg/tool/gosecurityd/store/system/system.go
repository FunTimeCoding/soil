package system

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/kernel_offer"
	"time"
)

type System struct {
	Identifier                 uint                  `gorm:"primarykey;column:id"`
	Name                       string                `gorm:"not null;uniqueIndex"`
	Release                    string                `gorm:"not null;default:''"`
	Kernel                     string                `gorm:"not null;default:''"`
	Priority                   int                   `gorm:"not null;default:0"`
	Labels                     map[string]string     `gorm:"serializer:json"`
	ListsUpdated               *time.Time            `gorm:"column:lists_updated"`
	KernelOffers               []*kernel_offer.Offer `gorm:"serializer:json"`
	KernelMetapackageInstalled *bool
	CreatedAt                  time.Time `gorm:"not null"`
	UpdatedAt                  time.Time `gorm:"not null"`
}
