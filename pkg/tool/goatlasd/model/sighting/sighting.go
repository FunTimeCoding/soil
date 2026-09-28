package sighting

import "github.com/funtimecoding/soil/pkg/tool/goatlasd/model/attribution"

type Sighting struct {
	Identifier              uint   `gorm:"primaryKey;column:identifier"`
	Source                  string `gorm:"column:source"`
	HardwareAddress         string `gorm:"column:hardware_address"`
	Address                 string `gorm:"column:address"`
	Hostname                string `gorm:"column:hostname"`
	Reserved                bool   `gorm:"column:reserved"`
	attribution.Attribution `gorm:"embedded"`
}
