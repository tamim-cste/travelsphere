function addToWishlist(countryName) {
    const feedback = document.getElementById('wishlist-feedback');
    feedback.innerHTML = '<span style="color:#888">Adding...</span>';

    fetch('/api/wishlist', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ country_name: countryName, status: 'Planned' })
    })
    .then(res => res.json())
    .then(data => {
        if (data.success) {
            feedback.innerHTML = '<span style="color:#6c63ff;font-weight:600">✓ Added to wishlist!</span>';
        } else {
            feedback.innerHTML = `<span style="color:#e74c3c">${data.message}</span>`;
        }
    })
    .catch(() => {
        feedback.innerHTML = '<span style="color:#e74c3c">Failed to add. Try again.</span>';
    });
}
