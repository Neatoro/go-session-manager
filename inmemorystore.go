package gosessionmanager

import (
	"errors"

	"github.com/google/uuid"
)

type storeInMemory struct {
	sessions map[string]Session
}

func NewInMemoryStore() SessionStore {
	return &storeInMemory{
		sessions: make(map[string]Session),
	}
}

func (store *storeInMemory) StartSession() (*Session, error) {
	session := Session{
		ID:   uuid.NewString(),
		Data: make(map[string]any),
	}
	store.sessions[session.ID] = session
	return &session, nil
}

func (store *storeInMemory) GetSession(id string) (*Session, error) {
	if session, ok := store.sessions[id]; ok {
		return &session, nil
	}
	return nil, errors.New("session not found")
}

func (store *storeInMemory) EndSession(session *Session) error {
	if _, ok := store.sessions[session.ID]; ok {
		delete(store.sessions, session.ID)
		return nil
	}
	return errors.New("session not found")
}
