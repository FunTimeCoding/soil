package store

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/proof_key"
)

func (s *Store) DeletePKCERequestSession(
	_ context.Context,
	signature string,
) error {
	return s.mapper.Where("signature = ?", signature).
		Delete(proof_key.Stub()).Error
}
