package da

import (
	"time"

	"code.google.com/p/go-uuid/uuid"
)

type Item struct {
	Id           string
	CreatedOn    time.Time
	Title        string
	Summary      string
	OwnersId     string
	ItemState    ItemState
	RecordStatus RecordStatus
	Position     int
}

func NewItem() *Item {
	return &Item{
		Id:           uuid.New(),
		CreatedOn:    time.Now(),
		ItemState:    Inception,
		RecordStatus: Active,
	}
}

func (e *Item) IsValid() bool {
	return e.Id != "" && e.Title != ""
}
