package claim

import "time"

type Claim struct {
	Identifier uint      `gorm:"primarykey;column:id" json:"id"`
	Item       string    `gorm:"not null;uniqueIndex:item_owner" json:"item"`
	Owner      string    `gorm:"not null;uniqueIndex:item_owner" json:"owner"`
	CreatedAt  time.Time `json:"created_at"`
}
