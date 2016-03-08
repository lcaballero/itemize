package da

type ItemState int

const (
	Inception  ItemState = 1
	InArchived ItemState = 2
	InProgress ItemState = 3
	InBacklog  ItemState = 4
	Completed  ItemState = 5
)

type RecordStatus int

const (
	Active  RecordStatus = 1
	Deleted RecordStatus = 2
)

type UserState int

const (
	Normal   UserState = 1
	Disabled UserState = 2
	Locked   UserState = 3
)
