<div class="auth-wrapper">
    <div class="auth-card">
        <div class="auth-logo">🌍</div>
        <h1>Welcome to TravelSphere</h1>
        <p class="auth-sub">Enter your username to access your wishlist and dashboard.</p>

        {{if .Error}}
        <div class="auth-error"><span>⚠️</span> {{.Error}}</div>
        {{end}}

        <form method="POST" action="/login">
            <div class="form-group">
                <label for="username">Username</label>
                <input type="text" id="username" name="username"
                       placeholder="e.g. beta" required autocomplete="username">
            </div>
            <button type="submit" class="btn-primary">Sign in →</button>
        </form>

        <p class="auth-hint">New user? Just enter any username to get started.</p>
    </div>
</div>
