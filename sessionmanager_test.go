package gosessionmanager

import (
	"context"
	"errors"
	"testing"
)

func TestReturnErrorIfThereIsNoActiveSession(t *testing.T) {
	sessionManager := NewInMemorySessionManager()
	err := sessionManager.EndSession(context.Background())
	if !errors.Is(err, ErrNoSession) {
		t.Error("TestReturnErrorIfThereIsNoActiveSession failed")
	}
}

func TestReturnErrorIfSessionDoesNotExist(t *testing.T) {
	sessionManager := NewInMemorySessionManager()
	err := sessionManager.EndSession(context.WithValue(context.Background(), sessionKey, "session-id"))
	if !errors.Is(err, ErrNoSession) {
		t.Error("TestReturnErrorIfSessionDoesNotExist failed")
	}
}
