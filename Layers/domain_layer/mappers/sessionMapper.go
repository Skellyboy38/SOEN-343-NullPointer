package mappers

import (
	"errors"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/classes"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/tdg"
)

type SessionMapper struct {
	sessions   map[int]classes.Session
	sessionTdg tdg.SessionTdg
}

func InitSessionMapper() *SessionMapper {
	return &SessionMapper{make(map[int]classes.Session), tdg.SessionTdg{}}
}

func (sessionMap *SessionMapper) InMemory(id int) bool {
	_, ok := sessionMap.sessions[id]
	if ok {
		return true
	} else {
		return false
	}
}

func (sessionMap *SessionMapper) Get(id int) (classes.Session, error) {
	if sessionMap.InMemory(id) {
		return sessionMap.sessions[id], nil
	} else {
		return classes.Session{}, errors.New("Session not in memory")
	}
}

func (sessionMap *SessionMapper) AddToMap(session classes.Session) {
	sessionMap.sessions[session.SessionId] = session
}
