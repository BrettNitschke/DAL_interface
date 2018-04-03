package cassandra

import (
	"time"

	"github.com/gocql/gocql"

	"git.sphere.ms/sphere/model/modelutil"
)

type Dal interface {
	GetUserByUuid(userId gocql.uuid) model.User
	//All other functions that access DB go here
}

type Session struct {
	*modelutil.Session
}

func GetSession() (*gocql.Session, error) {
	if session == nil {
		if err := modelutil.WaitFromEnv(time.Minute); err != nil {
			return nil, err
		}
	}
	if session == nil || session.Closed() {
		var err error
		session, err = modelutil.NewUpdatedKeyspaceSessionFromEnv()
		if err != nil {
			return nil, err
		}
	}
	return &Session{session.Session}, nil
}

func (session *Session) GetUserByUuid(userId gocql.uuid) model.User {
	//DB access stuff goes here
}
