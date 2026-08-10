package constant

const (
	SidebarFilterScript = `
	function filterSidebar(query) {
		var entries = document.querySelectorAll('.sidebar-entry');
		var lower = query.toLowerCase();
		entries.forEach(function(e) {
			var name = e.querySelector('.entry-name span');
			if (!name) return;
			e.style.display = name.textContent.toLowerCase().indexOf(lower) !== -1 ? '' : 'none';
		});
	}
`
	InfiniteScrollScript = `
	(function() {
		var sidebar = document.querySelector('.sidebar');
		if (!sidebar) return;
		var loading = false;
		sidebar.addEventListener('scroll', function() {
			if (loading) return;
			if (sidebar.scrollTop + sidebar.clientHeight < sidebar.scrollHeight - 50) return;
			var sentinel = sidebar.querySelector('[data-load-more]');
			if (!sentinel) return;
			loading = true;
			var url = sentinel.getAttribute('data-load-more');
			fetch(url).then(function(r) { return r.text(); }).then(function(html) {
				sentinel.outerHTML = html;
				loading = false;
			});
		});
	})();
`
	ScrollToBottomScript = `
	document.addEventListener('htmx:afterSwap', function(e) {
		if (e.detail.target.id === 'panel') {
			e.detail.target.scrollTop = e.detail.target.scrollHeight;
		}
	});
`
)
