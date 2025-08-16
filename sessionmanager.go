package gosessionmanager

import (
	"context"
	"errors"
)

type SessionManager struct {
	store SessionStore
}

type SessionStore interface {
	StartSession() (*Session, error)
	GetSession(id string) (*Session, error)
	EndSession(session *Session) error
}

type Session struct {
	ID   string
	Data map[string]any
}

func NewInMemorySessionManager() SessionManager {
	return SessionManager{
		store: NewInMemoryStore(),
	}
}

type sessionContextKeyType struct{}

var sessionKey = sessionContextKeyType{}

var ErrFailedStartingSession = errors.New("failed to start session")
var ErrNoSession = errors.New("no session found to end")

func (manager *SessionManager) StartSession(ctx context.Context) (context.Context, error) {
	if session, err := manager.store.StartSession(); err == nil {
		return context.WithValue(ctx, sessionKey, session.ID), nil
	}
	return nil, ErrFailedStartingSession
}

func (manager *SessionManager) EndSession(ctx context.Context) error {
	sessionId := ctx.Value(sessionKey)
	if sessionId == nil {
		return ErrNoSession
	}

	session, err := manager.store.GetSession(sessionId.(string))
	if err != nil {
		return ErrNoSession
	}

	err = manager.store.EndSession(session)
	return err
}
