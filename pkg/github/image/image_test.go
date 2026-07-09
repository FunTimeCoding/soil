package image

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/github/constant"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"github.com/google/go-github/v89/github"
	"testing"
	"time"
)

func TestImage(t *testing.T) {
	meta, e := json.Marshal(
		github.PackageMetadata{
			PackageType: new(constant.ContainerPackageType),
			Container: &github.PackageContainerMetadata{
				Tags: []string{upper.Alfa},
			},
		},
	)
	errors.PanicOnError(e)
	i := New(
		&github.PackageVersion{
			ID:        new(int64(1)),
			Name:      new(upper.Bravo),
			CreatedAt: &github.Timestamp{},
			Metadata:  meta,
		},
	)
	i.Raw = nil
	assert.Any(
		t,
		&Image{
			Identifier: 1,
			Digest:     upper.Bravo,
			Tags:       []string{upper.Alfa},
			Create:     time.Time{},
		},
		i,
	)
}
