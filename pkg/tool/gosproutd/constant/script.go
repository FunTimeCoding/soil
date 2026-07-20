package constant

const SortableScript = `
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
