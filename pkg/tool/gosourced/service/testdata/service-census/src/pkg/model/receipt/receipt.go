package receipt

import "example/pkg/constant"

type Receipt struct {
	Name string
}

func (Receipt) TableName() string {
	return constant.ReceiptsTable
}
