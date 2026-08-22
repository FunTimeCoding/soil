package service

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
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
		return validation.New("unknown relation type: %s", relationType)
	}

	return s.store.CreateRelation(
		sourceIdentifier,
		targetIdentifier,
		relationType,
	)
}
