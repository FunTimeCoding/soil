package constant

const InlineStyle = `
.status-dot {
	display: inline-block;
	width: 0.65rem;
	height: 0.65rem;
	border-radius: 50%;
	background: #5f5648;
}
.status-running { background: #4ab860; }
.floor-table td { padding: 0.5rem 0.75rem; }
.floor-table tr.stopped { opacity: 0.45; }
.floor-table .load { color: var(--pico-secondary); font-size: 0.85rem; }
.node-header {
	display: flex;
	align-items: baseline;
	gap: 0.75rem;
	margin: 1.25rem 0 0.5rem 0;
}
.node-header h3 { margin-bottom: 0; }
.node-header .release { color: var(--pico-muted-color); }
.badge {
	display: inline-block;
	padding: 0.1rem 0.55rem;
	border-radius: 0.35rem;
	font-size: 0.8rem;
}
.badge.updates { background: #3d2f14; color: #d8a03f; }
.badge.unbacked { background: #3d1a17; color: #e06a5e; }
`
