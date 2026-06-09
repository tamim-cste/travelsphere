<div class="page-header">
    <h1>Travel Dashboard</h1>
    <p>Your saved trips at a glance. Stats refresh automatically when your wishlist changes.</p>
</div>

<div id="dashboard-stats">
    <div class="stats-grid">
        <div class="stat-card">
            <span class="stat-label">TOTAL SAVED</span>
            <span class="stat-value">{{.Summary.Total}}</span>
        </div>
        <div class="stat-card">
            <span class="stat-label">PLANNED</span>
            <span class="stat-value">{{.Summary.Planned}}</span>
        </div>
        <div class="stat-card">
            <span class="stat-label">VISITED</span>
            <span class="stat-value">{{.Summary.Visited}}</span>
        </div>
    </div>
</div>

<div class="section">
    <h2>Saved destinations</h2>
    <div class="dest-list">
        {{range .Summary.Items}}
        <div class="dest-row">
            <span class="dest-name">{{.CountryName}}</span>
            <span class="dest-status {{.Status}}">{{.Status}}</span>
            {{if .Note}}<span class="dest-note">· {{.Note}}</span>{{end}}
        </div>
        {{else}}
        <p style="color:#888">No saved destinations yet.</p>
        {{end}}
    </div>
</div>

<script src="/static/js/dashboard.js"></script>
