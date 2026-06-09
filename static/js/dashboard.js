function refreshStats() {
    fetch('/api/dashboard/summary')
        .then(res => res.json())
        .then(data => {
            if (!data.success) return;
            const s = data.data;
            document.getElementById('dashboard-stats').innerHTML = `
                <div class="stats-grid">
                    <div class="stat-card">
                        <span class="stat-label">TOTAL SAVED</span>
                        <span class="stat-value">${s.total}</span>
                    </div>
                    <div class="stat-card">
                        <span class="stat-label">PLANNED</span>
                        <span class="stat-value">${s.planned}</span>
                    </div>
                    <div class="stat-card">
                        <span class="stat-label">VISITED</span>
                        <span class="stat-value">${s.visited}</span>
                    </div>
                </div>
            `;
        });
}
