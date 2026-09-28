package signing_key

import "time"

type SigningKey struct {
	Identifier string    `gorm:"primaryKey;column:identifier"`
	PrivateKey string    `gorm:"column:private_key"`
	PublicKey  string    `gorm:"column:public_key"`
	Algorithm  string    `gorm:"column:algorithm"`
	CreatedAt  time.Time `gorm:"column:created_at"`
}
