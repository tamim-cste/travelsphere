function saveRow(btn) {
    const row = btn.closest('tr');
    const id = row.dataset.id;
    const note = row.querySelector('.note-input').value;
    const status = row.querySelector('.status-select').value;

    fetch(`/api/wishlist/${id}`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ note, status })
    })
    .then(res => res.json())
    .then(data => {
        if (data.success) {
            btn.textContent = 'Saved ✓';
            setTimeout(() => btn.textContent = 'Save', 1500);
            refreshDashboardStats();
        }
    })
    .catch(() => alert('Save failed'));
}

function deleteRow(btn) {
    const row = btn.closest('tr');
    const id = row.dataset.id;

    fetch(`/api/wishlist/${id}`, { method: 'DELETE' })
    .then(res => res.json())
    .then(data => {
        if (data.success) {
            row.remove();
           
            const tbody = document.getElementById('wishlist-rows');
            if (tbody.children.length === 0) {
                tbody.innerHTML = '<tr><td colspan="4" style="text-align:center;color:#888;padding:2rem">No destinations yet.</td></tr>';
            }
            refreshDashboardStats();
        }
    })
    .catch(() => alert('Delete failed'));
}

function refreshDashboardStats() {
   
    fetch('/api/dashboard/summary').catch(() => {});
}
