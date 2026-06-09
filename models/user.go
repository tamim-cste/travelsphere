package models

// User represents an authenticated session user stored in memory
type User struct {
	Username string `json:"username"`
}
