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
	RecordStatus int
	UserState    int
}

func NewUser() *User {
	return &User{
		Id: uuid.New(),
	}
}

func (u *User) IsValid() bool {
	return u.FirstName != "" && u.Id != ""
}
