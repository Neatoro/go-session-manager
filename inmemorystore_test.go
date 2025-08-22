package gosessionmanager

import (
	"testing"
)

func TestStartSession(t *testing.T) {
	store := NewInMemoryStore()
	session, err := store.StartSession("some-session-id")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if session == nil {
		t.Fatal("expected session, got nil")
	}
	if session.ID != "some-session-id" {
		t.Errorf("expected session ID %v to be set to some-session-id", session.ID)
	}
	if session.Data == nil {
		t.Error("expected session data to be initialized")
	}
}

func TestGetSession(t *testing.T) {
	store := NewInMemoryStore()
	store.StartSession("some-session-id")
	got, err := store.GetSession("some-session-id")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got == nil || got.ID != "some-session-id" {
		t.Errorf("expected session ID %v, got %v", "some-session-id", got)
	}
}

func TestGetSessionShouldReturnErrorIfNotExisting(t *testing.T) {
	store := NewInMemoryStore()

	if _, err := store.GetSession("non-existent-id"); err != ErrNoSession {
		t.Errorf("expected ErrNoSession, got %v", err)
	}
}

func TestEndSession(t *testing.T) {
	store := NewInMemoryStore()
	session, _ := store.StartSession("some-session-id")
	err := store.EndSession(session)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = store.EndSession(session)
	if err != ErrNoSession {
		t.Errorf("expected ErrNoSession, got %v", err)
	}
}

func TestUpdateSessionReturnsNil(t *testing.T) {
	store := NewInMemoryStore()
	session, _ := store.StartSession("some-session-id")
	err := store.UpdateSession(session)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
