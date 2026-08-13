package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"slices"
)

func (s *Service) CreateRelation(
	sourceIdentifier int64,
	targetIdentifier int64,
	relationType string,
) error {
	if relationType != "" &&
		!slices.Contains(constant.RelationTypes, relationType) {
		return fmt.Errorf("%w: %s", constant.ErrorRelationType, relationType)
	}

	return s.store.CreateRelation(
		sourceIdentifier,
		targetIdentifier,
		relationType,
	)
}
