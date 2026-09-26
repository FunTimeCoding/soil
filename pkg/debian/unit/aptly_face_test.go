package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/debian/aptly"
	"github.com/funtimecoding/soil/pkg/debian/aptly/face"
	"github.com/funtimecoding/soil/pkg/debian/aptly/mock_client"
	"testing"
)

func TestAptlyClientSatisfiesTheRepositoryFace(t *testing.T) {
	var v face.Repository = aptly.New(
		"host.example",
		443,
		false,
		"user",
		"password",
	)
	assert.NotNil(t, v)
}

func TestAptlyMockSatisfiesTheRepositoryFace(t *testing.T) {
	var v face.Repository = mock_client.New()
	assert.NotNil(t, v)
}

func TestAptlyMockServesSeededVersions(t *testing.T) {
	c := mock_client.New()
	c.SeedVersions("stable", "foxtrot", "1.2.5", "1.2.3", "1.2.0")
	v, e := c.Versions("stable", "foxtrot")
	assert.Nil(t, e)
	assert.Count(t, 3, v)
	assert.String(t, "1.2.5", v[0])
}

func TestAptlyMockLatestVersionTakesTheFirst(t *testing.T) {
	c := mock_client.New()
	c.SeedVersions("stable", "foxtrot", "1.2.5", "1.2.3")
	v, e := c.LatestVersion("stable", "foxtrot")
	assert.Nil(t, e)
	assert.String(t, "1.2.5", v)
}

func TestAptlyMockReportsAnAbsentPackage(t *testing.T) {
	c := mock_client.New()
	v, e := c.LatestVersion("stable", "absent")
	assert.Nil(t, e)
	assert.String(t, "", v)
}

func TestAptlyMockPropagatesFailure(t *testing.T) {
	c := mock_client.New()
	c.Fail(errors.New("repository unreachable"))
	_, e := c.Versions("stable", "foxtrot")
	assert.NotNil(t, e)
}
