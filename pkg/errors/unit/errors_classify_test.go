package unit

import (
	"errors"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unit/expiring"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"testing"
)

func TestClassifyReportsFaults(t *testing.T) {
	assert.True(t, classify.Reportable(errors.New("connection reset")))
	assert.True(t, classify.Reportable(not_found.Format("post not found: 8")))
}

func TestClassifyKeepsCallerErrorsOut(t *testing.T) {
	assert.False(t, classify.Reportable(validation.New("page 9 out of range")))
	assert.False(
		t,
		classify.Reportable(fmt.Errorf("refresh: %w", expiring.New())),
	)
}

func TestClassifyMessageCarriesTypedDetail(t *testing.T) {
	assert.String(
		t,
		"post failed: post not found: 8",
		classify.Message(not_found.Format("post not found: 8"), "post failed"),
	)
	assert.String(
		t,
		"post failed",
		classify.Message(errors.New("connection reset"), "post failed"),
	)
}
