package mappers

import (
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/classes"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/tdg"
)

type SessionMapper struct {
	sessions        map[int]classes.Session
	usersToSessions map[int]classes.Session
	SessionTdg      tdg.SessionTdg
}

func InitSessionMapper() *SessionMapper {
	return &SessionMapper{make(map[int]classes.Session), make(map[int]classes.Session), tdg.SessionTdg{}}
}

func (sessionMap *SessionMapper) InMemoryByUser(user classes.User) bool {
	_, ok := sessionMap.usersToSessions[user.StudentId]
	if ok {
		return true
	} else {
		return false
	}
}

func (sessionMap *SessionMapper) InMemoryBySessionId(sessionId int) bool {
	_, ok := sessionMap.sessions[sessionId]
	if ok {
		return true
	} else {
		return false
	}
}

func (sessionMap *SessionMapper) Get(user classes.User) (classes.Session, error) {
	if sessionMap.InMemoryByUser(user) {
		return sessionMap.usersToSessions[user.StudentId], nil
	} else {
		sessionId, _, err := MapperBundle.SessionMapper.SessionTdg.Read(user.StudentId)
		if err == nil {
			student := MapperBundle.UserMapper.users[user.StudentId]
			currentSession := classes.Session{sessionId, student}
			sessionMap.AddToMap(currentSession)
			return currentSession, nil
		}
		sessionId, err = createSession(user.StudentId)
		student := MapperBundle.UserMapper.users[user.StudentId]
		currentSession := classes.Session{sessionId, student}
		sessionMap.AddToMap(currentSession)

		return currentSession, nil
	}
}

func (sessionMap *SessionMapper) AddToMap(session classes.Session) {
	sessionMap.usersToSessions[session.User.StudentId] = session
	sessionMap.sessions[session.SessionId] = session
}

func createSession(studentId int) (int, error) {
	return tdg.SessionTdg{}.Create(studentId)
}
