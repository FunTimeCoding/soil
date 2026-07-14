package web

const inlineCSS = `
.seed-table { width: 100%; }
.seed-table th, .seed-table td { padding: 0.5rem 0.75rem; }
.position-cell { width: 3rem; text-align: center; }
.seed-table tbody.sortable tr { cursor: grab; }
.seed-table tbody.sortable tr:active { cursor: grabbing; }
.sortable-ghost { opacity: 0.4; background: var(--pico-primary-focus); }
.sortable-chosen { background: var(--pico-card-background-color); }
.drag-handle { cursor: grab; user-select: none; color: var(--pico-muted-color); }
`

const sortableJS = `
htmx.onLoad(function(content) {
	var sortables = content.querySelectorAll(".sortable");
	for (var i = 0; i < sortables.length; i++) {
		var sortable = sortables[i];
		var instance = new Sortable(sortable, {
			animation: 150,
			ghostClass: 'sortable-ghost',
			chosenClass: 'sortable-chosen',
			handle: '.drag-handle',
			onEnd: function (evt) {
				var items = evt.from.querySelectorAll('input[name="item"]');
				var ids = Array.from(items).map(function(i) { return i.value; });
				htmx.ajax('POST', '/reorder', {
					values: { item: ids },
					swap: 'none'
				});
			}
		});
	}
});
`
