package subscription

import "time"

type Subscription struct {
	Identifier uint      `gorm:"primaryKey;autoIncrement;column:identifier"`
	Callsign   string    `gorm:"index;column:callsign"`
	Instance   string    `gorm:"index:subscription_request;column:instance"`
	Project    int64     `gorm:"index:subscription_request;column:project"`
	Request    int64     `gorm:"index:subscription_request;column:request"`
	Reference  string    `gorm:"column:reference"`
	Alias      string    `gorm:"column:alias"`
	Delivery   string    `gorm:"column:delivery"`
	CreatedAt  time.Time `gorm:"column:created_at"`
}
