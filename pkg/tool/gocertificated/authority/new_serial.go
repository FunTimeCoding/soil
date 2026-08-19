package authority

import (
	"crypto/rand"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"math/big"
)

func newSerial() *big.Int {
	s, e := rand.Int(
		rand.Reader,
		new(big.Int).Lsh(big.NewInt(1), constant.SerialBit),
	)
	errors.PanicOnError(e)

	return s
}
