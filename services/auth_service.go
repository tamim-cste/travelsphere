package services

import (
	"fmt"
	"sync"
	"travelsphere/models"
)

// userStore holds registered usernames in memory
var (
	userStore = map[string]models.User{
		"beta": {Username: "beta"},
	}
	userMu sync.RWMutex
)

// GetOrCreateUser returns an existing user or registers a new username
func GetOrCreateUser(username string) (*models.User, error) {
	if username == "" {
		return nil, fmt.Errorf("username is required")
	}

	userMu.Lock()
	defer userMu.Unlock()

	if user, ok := userStore[username]; ok {
		return &user, nil
	}

	// Register new username on first login
	newUser := models.User{Username: username}
	userStore[username] = newUser
	return &newUser, nil
}

// User Exists checks if a username is registered
func UserExists(username string) bool {
	userMu.RLock()
	defer userMu.RUnlock()
	_, ok := userStore[username]
	return ok
}
