package constant

const FixtureSilences = `[{
	"id": "Alfa",
	"status": {"state": "active"},
	"createdBy": "Bravo",
	"comment": "Charlie",
	"startsAt": "2026-01-01T00:00:00Z",
	"endsAt": "2026-01-01T01:00:00Z",
	"updatedAt": "2026-01-01T00:00:00Z",
	"matchers": []
}]`

const FixtureSuppressedAndUnratedAlerts = `[
	{
		"annotations": {},
		"endsAt": "2026-01-01T02:00:00Z",
		"fingerprint": "alfa",
		"receivers": [{"name": "Charlie"}],
		"startsAt": "2026-01-01T00:00:00Z",
		"status": {"inhibitedBy": [], "silencedBy": ["Delta"], "state": "suppressed"},
		"updatedAt": "2026-01-01T00:00:00Z",
		"labels": {"alertname": "Alfa", "severity": "critical"}
	},
	{
		"annotations": {},
		"endsAt": "2026-01-01T02:00:00Z",
		"fingerprint": "bravo",
		"receivers": [{"name": "Charlie"}],
		"startsAt": "2026-01-01T00:00:00Z",
		"status": {"inhibitedBy": [], "silencedBy": [], "state": "active"},
		"updatedAt": "2026-01-01T00:00:00Z",
		"labels": {"alertname": "Bravo", "severity": "none"}
	}
]`
