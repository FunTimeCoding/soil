package integration

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/unit/terraform_tester"
	"testing"
	"time"
)

func TestLockFromGonePodIsCleared(t *testing.T) {
	c := mock_client.New().WithLease(
		terraform_tester.HeldLease(constant.FixtureGoneHolder, time.Now()),
	)
	o := terraform_tester.New(t, c)
	assert.Strings(
		t,
		[]string{
			"init -json",
			"force-unlock -force 11111111-2222-4333-8444-666666666666",
			"plan -out=PLAN -no-color",
			"show -json PLAN",
			"apply -auto-approve -no-color PLAN",
		},
		o.Calls(),
	)
	assert.Strings(t, []string{"terraform/runner-5c7d9e8f1-fghij"}, c.Lookups())
	assert.Strings(
		t,
		[]string{"terraform state lock orphaned, force-unlocked"},
		o.Messages(),
	)
	assert.Any(
		t,
		"root@runner-5c7d9e8f1-fghij",
		o.Reporter.Events()[0].Context["who"],
	)
}

func TestLockFromLivePodIsKept(t *testing.T) {
	c := mock_client.New().WithLease(
		terraform_tester.HeldLease(constant.FixtureGoneHolder, time.Now()),
	).WithPod(constant.FixtureGonePod)
	o := terraform_tester.New(t, c)
	assert.Strings(
		t,
		[]string{
			"init -json",
			"plan -out=PLAN -no-color",
			"show -json PLAN",
			"apply -auto-approve -no-color PLAN",
		},
		o.Calls(),
	)
	assert.Count(t, 0, o.Messages())
}

func TestLockFromOutsideTheClusterIsKept(t *testing.T) {
	c := mock_client.New().WithLease(
		terraform_tester.HeldLease(
			constant.FixtureLaptop,
			time.Now().Add(-time.Hour),
		),
	)
	o := terraform_tester.New(t, c)
	assert.Strings(
		t,
		[]string{
			"init -json",
			"plan -out=PLAN -no-color",
			"show -json PLAN",
			"apply -auto-approve -no-color PLAN",
		},
		o.Calls(),
	)
	assert.Count(t, 0, c.Lookups())
}

func TestLockFromAnotherDeploymentIsKept(t *testing.T) {
	c := mock_client.New().WithLease(
		terraform_tester.HeldLease(constant.FixtureOtherHolder, time.Now()),
	)
	o := terraform_tester.New(t, c)
	assert.Strings(
		t,
		[]string{
			"init -json",
			"plan -out=PLAN -no-color",
			"show -json PLAN",
			"apply -auto-approve -no-color PLAN",
		},
		o.Calls(),
	)
	assert.Count(t, 0, c.Lookups())
}

func TestOwnLockFromBeforeStartIsCleared(t *testing.T) {
	c := mock_client.New().WithLease(
		terraform_tester.HeldLease(
			constant.FixtureOwnHolder,
			time.Now().Add(-time.Hour),
		),
	)
	o := terraform_tester.New(t, c)
	assert.Strings(
		t,
		[]string{
			"init -json",
			"force-unlock -force 11111111-2222-4333-8444-666666666666",
			"plan -out=PLAN -no-color",
			"show -json PLAN",
			"apply -auto-approve -no-color PLAN",
		},
		o.Calls(),
	)
	assert.Count(t, 0, c.Lookups())
}

func TestOwnLockFromAfterStartIsKept(t *testing.T) {
	c := mock_client.New().WithLease(
		terraform_tester.HeldLease(
			constant.FixtureOwnHolder,
			time.Now().Add(time.Hour),
		),
	)
	o := terraform_tester.New(t, c)
	assert.Strings(
		t,
		[]string{
			"init -json",
			"plan -out=PLAN -no-color",
			"show -json PLAN",
			"apply -auto-approve -no-color PLAN",
		},
		o.Calls(),
	)
	assert.Count(t, 0, c.Lookups())
}

func TestUnknownHolderIsReportedAndKept(t *testing.T) {
	c := mock_client.New().WithLease(
		terraform_tester.HeldLease(constant.FixtureGoneHolder, time.Now()),
	).WithPodError(fmt.Errorf("pods is forbidden"))
	o := terraform_tester.New(t, c)
	assert.Strings(
		t,
		[]string{
			"init -json",
			"plan -out=PLAN -no-color",
			"show -json PLAN",
			"apply -auto-approve -no-color PLAN",
		},
		o.Calls(),
	)
	assert.Strings(t, []string{"pods is forbidden"}, o.Messages())
}

func TestTruncatedHolderIsKept(t *testing.T) {
	assert.Integer(t, 63, len(constant.FixtureLongOther))
	c := mock_client.New().WithLease(
		terraform_tester.HeldLease(constant.FixtureLongOtherHolder, time.Now()),
	)
	o := terraform_tester.NewAs(t, c, constant.FixtureLongOwn)
	assert.Strings(
		t,
		[]string{
			"init -json",
			"plan -out=PLAN -no-color",
			"show -json PLAN",
			"apply -auto-approve -no-color PLAN",
		},
		o.Calls(),
	)
	assert.Count(t, 0, c.Lookups())
}

func TestFreeLockSkipsUnlock(t *testing.T) {
	o := terraform_tester.New(t, mock_client.New())
	assert.Strings(
		t,
		[]string{
			"init -json",
			"plan -out=PLAN -no-color",
			"show -json PLAN",
			"apply -auto-approve -no-color PLAN",
		},
		o.Calls(),
	)
	assert.Count(t, 0, o.Messages())
}

func TestFailedUnlockIsReported(t *testing.T) {
	c := mock_client.New().WithLease(
		terraform_tester.HeldLease(constant.FixtureGoneHolder, time.Now()),
	)
	o := terraform_tester.NewFailingUnlock(t, c)
	assert.Strings(
		t,
		[]string{
			"init -json",
			"force-unlock -force 11111111-2222-4333-8444-666666666666",
			"plan -out=PLAN -no-color",
			"show -json PLAN",
			"apply -auto-approve -no-color PLAN",
		},
		o.Calls(),
	)
	m := o.Messages()
	assert.Count(t, 1, m)
	assert.Suffix(
		t,
		"recorder force-unlock -force 11111111-2222-4333-8444-666666666666: exit status 1",
		m[0],
	)
}

func TestReleasedLeaseSkipsUnlock(t *testing.T) {
	c := mock_client.New().WithLease(
		terraform_tester.ReleasedLease(constant.FixtureGoneHolder, time.Now()),
	)
	o := terraform_tester.New(t, c)
	assert.Strings(
		t,
		[]string{
			"init -json",
			"plan -out=PLAN -no-color",
			"show -json PLAN",
			"apply -auto-approve -no-color PLAN",
		},
		o.Calls(),
	)
	assert.Count(t, 0, c.Lookups())
}
