package mappers

import (
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/classes"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/tdg"
)

type SessionMapper struct {
	sessions   map[int]classes.Session
	SessionTdg tdg.SessionTdg
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
		sessionId, _, err := MapperBundle.SessionMapper.SessionTdg.Read(id)
		if err == nil{
			student := MapperBundle.UserMapper.users[id]
			currentSession := classes.Session{sessionId,student}
			sessionMap.AddToMap(currentSession)
			return currentSession, nil
		}
		sessionId, err = createSession(id)
		student := MapperBundle.UserMapper.users[id]
		currentSession := classes.Session{sessionId,student}
		sessionMap.AddToMap(currentSession)

		return currentSession, nil
	}
}

func (sessionMap *SessionMapper) AddToMap(session classes.Session) {
	sessionMap.sessions[session.User.StudentId] = session
}

func createSession(studentId int) (int, error){
	return tdg.SessionTdg{}.Create(studentId)
}