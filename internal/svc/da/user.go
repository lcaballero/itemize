package da

import (
	"time"

	"code.google.com/p/go-uuid/uuid"
)

type User struct {
	Id           string
	CreatedOn    time.Time
	Username     string
	FirstName    string
	LastName     string
	Email        string
	Icon         string
	RecordStatus RecordStatus
	UserState    UserState
}

func NewUser() User {
	return User{
		Id:           uuid.New(),
		CreatedOn:    time.Now(),
		RecordStatus: Active,
		UserState:    Normal,
	}
}

func (u User) IsValid() bool {
	return u.FirstName != "" && u.Id != ""
}
