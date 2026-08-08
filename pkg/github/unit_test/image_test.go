package unit_test

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/github/constant"
	"github.com/funtimecoding/soil/pkg/github/image"
	strings "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/google/go-github/v90/github"
	"testing"
	"time"
)

func TestImage(t *testing.T) {
	meta, e := json.Marshal(
		github.PackageMetadata{
			PackageType: new(constant.ContainerPackageType),
			Container: &github.PackageContainerMetadata{
				Tags: []string{strings.UpperAlfa},
			},
		},
	)
	errors.PanicOnError(e)
	i := image.New(
		&github.PackageVersion{
			ID:        new(int64(1)),
			Name:      new(strings.UpperBravo),
			CreatedAt: &github.Timestamp{},
			Metadata:  meta,
		},
	)
	i.Raw = nil
	assert.Any(
		t,
		&image.Image{
			Identifier: 1,
			Digest:     "Bravo",
			Tags:       []string{"Alfa"},
			Create:     time.Time{},
		},
		i,
	)
}
