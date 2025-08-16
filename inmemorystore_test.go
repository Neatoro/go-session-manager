package gosessionmanager

import (
	"math/rand"
	"testing"

	"github.com/google/uuid"
)

func TestStartSession(t *testing.T) {
	uuid.SetRand(rand.New(rand.NewSource(1)))

	store := NewInMemoryStore()
	session, err := store.StartSession()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if session == nil {
		t.Fatal("expected session, got nil")
	}
	if session.ID != "52fdfc07-2182-454f-963f-5f0f9a621d72" {
		t.Errorf("expected session ID %v to be set to 52fdfc07-2182-454f-963f-5f0f9a621d72", session.ID)
	}
	if session.Data == nil {
		t.Error("expected session data to be initialized")
	}
}

func TestGetSession(t *testing.T) {
	store := NewInMemoryStore()
	session, _ := store.StartSession()
	got, err := store.GetSession(session.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got == nil || got.ID != session.ID {
		t.Errorf("expected session ID %v, got %v", session.ID, got)
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
	session, _ := store.StartSession()
	err := store.EndSession(session)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = store.EndSession(session)
	if err != ErrNoSession {
		t.Errorf("expected ErrNoSession, got %v", err)
	}
}
