package gosessionmanager

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
