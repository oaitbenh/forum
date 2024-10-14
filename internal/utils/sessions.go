package web_forum

import (
	"database/sql"
	"fmt"
	"html/template"
	"sync"
	"time"

	"github.com/google/uuid"
)

var (
	Db        *sql.DB
	S_essions *SessionManager
	Template  *template.Template
)

type Session struct {
	Username string
	Expiry   time.Time
}

type SessionManager struct {
	sessions map[string]Session
	mutex    sync.RWMutex
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessions: make(map[string]Session),
	}
}

func (sm *SessionManager) CreateSession(username string) string {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	sessionID := uuid.New().String()
	expiry := time.Now().Add(30 * time.Minute)
	sm.sessions[sessionID] = Session{
		Username: username,
		Expiry:   expiry,
	}
	return sessionID
}

func (sm *SessionManager) GetSession(sessionID string) (Session, bool) {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()
	session, found := sm.sessions[sessionID]
	if !found || time.Now().After(session.Expiry) {
		delete(sm.sessions, sessionID)
		// sm.DeleteSession(sessionID)
		return Session{}, false
	}
	return session, true
}

func (sm *SessionManager) DeleteSession(sessionID string) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	delete(sm.sessions, sessionID)
	fmt.Println("1", sessionID)
}
