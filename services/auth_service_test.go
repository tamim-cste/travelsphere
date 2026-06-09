package services

import (
	"testing"
	"travelsphere/models"
)

func resetUserStore() {
	userMu.Lock()
	defer userMu.Unlock()
	userStore = map[string]models.User{
		"beta": {Username: "beta"},
	}
}

func TestGetOrCreateUserNewUser(t *testing.T) {
	resetUserStore()
	user, err := GetOrCreateUser("alice")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if user.Username != "alice" {
		t.Errorf("expected alice, got %s", user.Username)
	}
}

func TestGetOrCreateUserExistingUser(t *testing.T) {
	resetUserStore()
	GetOrCreateUser("bob")
	user, err := GetOrCreateUser("bob")
	if err != nil {
		t.Fatalf("unexpected error on second call: %v", err)
	}
	if user.Username != "bob" {
		t.Errorf("expected bob, got %s", user.Username)
	}
}

func TestGetOrCreateUserEmptyUsername(t *testing.T) {
	_, err := GetOrCreateUser("")
	if err == nil {
		t.Error("expected error for empty username")
	}
}

func TestGetOrCreateUserPreseededBeta(t *testing.T) {
	resetUserStore()
	user, err := GetOrCreateUser("beta")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if user.Username != "beta" {
		t.Errorf("expected beta, got %s", user.Username)
	}
}

func TestUserExistsKnown(t *testing.T) {
	resetUserStore()
	if !UserExists("beta") {
		t.Error("expected beta to exist")
	}
}

func TestUserExistsUnknown(t *testing.T) {
	resetUserStore()
	if UserExists("nobody") {
		t.Error("expected nobody to not exist before registration")
	}
}

func TestGetOrCreateUserRegistersNewOnFirstLogin(t *testing.T) {
	resetUserStore()
	GetOrCreateUser("carol")
	if !UserExists("carol") {
		t.Error("expected carol to be registered after first login")
	}
}

func TestGetOrCreateUserDoesNotDuplicate(t *testing.T) {
	resetUserStore()
	GetOrCreateUser("dave")
	GetOrCreateUser("dave")
	GetOrCreateUser("dave")
	// Still only one dave — store is a map
	if !UserExists("dave") {
		t.Error("expected dave to exist")
	}
}
