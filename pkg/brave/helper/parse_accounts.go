package helper

import (
	"github.com/funtimecoding/soil/pkg/brave/constant"
	"google.golang.org/protobuf/encoding/protowire"
)

func ParseAccounts(b []byte) (name string, email string, photo string, gaia string) {
	position := 0

	for position < len(b) {
		fieldNum, wireType, n := protowire.ConsumeTag(b[position:])

		if n < 0 {
			break
		}

		position += n

		switch wireType {
		case protowire.BytesType:
			v, y := protowire.ConsumeBytes(b[position:])

			if y < 0 {
				break
			}

			position += y

			if fieldNum == constant.AccountNestedField {
				return ParseAccounts(v)
			}

			switch fieldNum {
			case constant.AccountNameField:
				name = string(v)
			case constant.AccountEmailField:
				email = string(v)
			case constant.AccountPhotoField:
				photo = string(v)
			case constant.AccountGaiaField:
				gaia = string(v)
			}
		case protowire.VarintType:
			_, y := protowire.ConsumeVarint(b[position:])

			if y < 0 {
				break
			}

			position += y
		default:
			y := protowire.ConsumeFieldValue(fieldNum, wireType, b[position:])

			if y < 0 {
				break
			}

			position += y
		}
	}

	return
}
