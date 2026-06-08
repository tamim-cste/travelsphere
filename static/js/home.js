function searchDestinations() {
    const query = document.getElementById('search-input').value;
    if (!query) return;

    // Loading state
    document.getElementById('search-suggestions').innerHTML = 
        '<p>Searching...</p>';

    fetch(`/api/countries?search=${encodeURIComponent(query)}`)
        .then(res => res.json())
        .then(data => {
            if (!data.success) {
                document.getElementById('search-suggestions').innerHTML = 
                    '<p>No results found</p>';
                return;
            }
            renderSuggestions(data.data);
        })
        .catch(() => {
            document.getElementById('search-suggestions').innerHTML = 
                '<p>Something went wrong</p>';
        });
}

function renderSuggestions(countries) {
 
}