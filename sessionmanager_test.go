package gosessionmanager

import (
	"context"
	"errors"
	"strconv"
	"testing"
)

type StoreStub struct {
	stubedSessions map[string]Session
}

func (s *StoreStub) EndSession(session *Session) error {
	delete(s.stubedSessions, session.ID)
	return nil
}

func (s *StoreStub) GetSession(id string) (*Session, error) {
	if session, ok := s.stubedSessions[id]; ok {
		return &session, nil
	}
	return nil, ErrNoSession
}

func (s *StoreStub) StartSession() (*Session, error) {
	sessionCount := len(s.stubedSessions)
	session := Session{
		ID:   "session-id-" + strconv.FormatInt(int64(sessionCount), 10),
		Data: make(map[string]any),
	}

	s.stubedSessions[session.ID] = session

	return &session, nil
}

func newStoreStub(stubedSessions map[string]Session) SessionStore {
	return &StoreStub{stubedSessions}
}

func TestNewInMemorySessionManager(t *testing.T) {
	manager := NewInMemorySessionManager()
	if manager.store == nil {
		t.Error("Expected store to be initialized, got nil")
	}
}

func TestShouldStartANewSession(t *testing.T) {
	store := newStoreStub(map[string]Session{})
	sessionManager := SessionManager{store}

	ctx := context.Background()
	ctx, err := sessionManager.StartSession(ctx)
	if err != nil {
		t.Error("TestShouldStartANewSession failed: failed to start session")
	}

	if _, err := store.GetSession(ctx.Value(sessionKey).(string)); err != nil {
		t.Error("TestShouldStartANewSession failed: found no session")
	}
}

func TestShouldEndSession(t *testing.T) {
	ctx := context.WithValue(context.Background(), sessionKey, "session-id-1")
	store := newStoreStub(map[string]Session{
		"session-id-1": Session{
			ID:   "session-id-1",
			Data: map[string]any{},
		},
	})
	sessionManager := SessionManager{store}

	if err := sessionManager.EndSession(ctx); err != nil {
		t.Error("TestShouldEndSession failed: failed to stop session")
	}

	if _, err := store.GetSession(ctx.Value(sessionKey).(string)); err == nil {
		t.Error("TestShouldEndSession failed: found session")
	}
}

func TestReturnErrorIfThereIsNoActiveSession(t *testing.T) {
	sessionManager := SessionManager{store: newStoreStub(map[string]Session{})}
	err := sessionManager.EndSession(context.Background())
	if !errors.Is(err, ErrNoSession) {
		t.Error("TestReturnErrorIfThereIsNoActiveSession failed")
	}
}

func TestReturnErrorIfSessionDoesNotExist(t *testing.T) {
	sessionManager := SessionManager{store: newStoreStub(map[string]Session{})}
	err := sessionManager.EndSession(context.WithValue(context.Background(), sessionKey, "session-id-1"))
	if !errors.Is(err, ErrNoSession) {
		t.Error("TestReturnErrorIfSessionDoesNotExist failed")
	}
}
