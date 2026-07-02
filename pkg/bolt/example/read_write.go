package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/bolt"
	"github.com/funtimecoding/go-library/pkg/bolt/constant"
	"go.etcd.io/bbolt"
)

func ReadWrite() {
	b := bolt.New("tmp/bolt.db")
	defer b.Close()
	b.MustUpdate(
		func(t *bbolt.Tx) error {
			b.MustPut(
				b.MustCreateBucket(t, constant.Bucket),
				constant.Key,
				"value",
			)

			return nil
		},
	)
	b.MustView(
		func(t *bbolt.Tx) error {
			v := b.Bucket(t, constant.Bucket)

			if v == nil {
				return nil
			}

			fmt.Println(b.Get(v, constant.Key))

			return nil
		},
	)
}
