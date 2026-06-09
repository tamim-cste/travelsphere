<div class="auth-wrapper">
    <div class="auth-card">
        <h1>Sign in</h1>
        <p class="auth-sub">Use <strong>beta / 1234</strong> to log in.</p>

        {{if .Error}}
        <p class="auth-error">{{.Error}}</p>
        {{end}}

        <form method="POST" action="/login">
            <div class="form-group">
                <label>Username</label>
                <input type="text" name="username" placeholder="beta" required>
            </div>
            <div class="form-group">
                <label>Password</label>
                <input type="password" name="password" placeholder="••••" required>
            </div>
            <button type="submit" class="btn-primary">Login</button>
        </form>
    </div>
</div>
