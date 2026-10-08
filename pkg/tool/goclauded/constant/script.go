package constant

const (
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
		if (e.detail.target.id !== 'panel') return;
		if (e.detail.target.querySelector('.search-current')) return;
		e.detail.target.scrollTop = e.detail.target.scrollHeight;
	});
`
	HitNavigationScript = `
	(function() {
		function hits() {
			return Array.prototype.slice.call(document.querySelectorAll('#panel .search-hit'));
		}
		function current() {
			return hits().findIndex(function(e) { return e.classList.contains('search-current'); });
		}
		function show(index) {
			var all = hits();
			if (!all.length) return;
			index = (index + all.length) % all.length;
			all.forEach(function(e) { e.classList.remove('search-current'); });
			all[index].classList.add('search-current');
			all[index].scrollIntoView({block: 'center'});
			var counter = document.getElementById('hit-counter');
			if (counter) counter.textContent = 'hit ' + (index + 1) + ' of ' + all.length + ' · n / p';
		}
		document.addEventListener('keydown', function(e) {
			if (e.target.tagName === 'INPUT') return;
			if (e.key === 'n') show(current() + 1);
			if (e.key === 'p') show(current() - 1);
		});
		document.addEventListener('htmx:afterSwap', function(e) {
			if (e.detail.target.id !== 'panel') return;
			var index = current();
			if (index >= 0) show(index);
		});
	})();
`
)
