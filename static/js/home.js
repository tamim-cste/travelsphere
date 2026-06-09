function searchDestinations() {
    const query = document.getElementById('search-input').value.trim();
    const box = document.getElementById('search-suggestions');

    if (!query) {
        box.innerHTML = '';
        return;
    }

    box.innerHTML = '<div class="suggestion-item">Searching...</div>';

    fetch(`/api/countries?search=${encodeURIComponent(query)}`)
        .then(res => res.json())
        .then(data => {
            if (!data.success || !data.data || data.data.length === 0) {
                box.innerHTML = '<div class="suggestion-item">No results found</div>';
                return;
            }
            renderSuggestions(data.data.slice(0, 8));
        })
        .catch(() => {
            box.innerHTML = '<div class="suggestion-item">Something went wrong</div>';
        });
}

function renderSuggestions(countries) {
    const box = document.getElementById('search-suggestions');
    box.innerHTML = countries.map(c => `
        <div class="suggestion-item" onclick="window.location.href='/countries/${c.slug}'">
            ${c.name} &mdash; ${c.capital || 'N/A'}
        </div>
    `).join('');
}

// Close suggestions when clicking outside
document.addEventListener('click', function(e) {
    const box = document.getElementById('search-suggestions');
    if (!e.target.closest('.search-wrapper')) {
        box.innerHTML = '';
    }
});
