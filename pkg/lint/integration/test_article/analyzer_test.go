package test_article

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"testing"
)

func TestArticleFlagsArticleInsideTestName(t *testing.T) {
	results := analyze(t)
	testutil.AssertBlocked(t, results, 4)

	for _, text := range []string{
		`TestAQueueDrainsOnClose carries the article "a"`,
		`TestTheReconcilerOnlyPromotes carries the article "the"`,
		`TestRetryAfterADelaySucceeds carries the article "a"`,
		`TestReadsAnHourOfLogs carries the article "an"`,
	} {
		testutil.AssertBlockedContains(t, results, text)
	}
}

func TestArticleSparesLookalikesLabelsAndNonTests(t *testing.T) {
	results := analyze(t)

	for _, name := range []string{
		"TestThemeLoads ",
		"TestAPIAnswers ",
		"TestDoorA ",
		"TestAnswerHolds ",
		"BenchmarkTheQueue ",
		"helperTheQueue ",
		"TestTheOutsideOfTestFiles ",
	} {
		testutil.AssertNotBlockedContains(t, results, name)
	}
}
