package da

import (
	"errors"
	"log"
	"os"
	"sync"
	"time"

	"github.com/lcaballero/hitman"
)

const DefaultDbName = "items.db.json"

type access func(*DataStore)

type AccessClient struct {
	store  *DataStore
	dbname string
	lock   *sync.Mutex
}

func NewAccessClient(dbname string) (*AccessClient, error) {
	var data *DataStore
	_, err := os.Stat(dbname)
	if os.IsNotExist(err) {
		data, err = NewDefaultDataStore()
	} else {
		data, err = LoadFromFile(dbname)
	}
	if err != nil {
		return nil, err
	}
	a := &AccessClient{
		store:  data,
		dbname: dbname,
		lock:   &sync.Mutex{},
	}
	return a, nil
}

func (d *AccessClient) Start() hitman.KillChannel {
	done := hitman.NewKillChannel()
	writeTic := time.NewTicker(5 * time.Second).C

	go func() {
		for {
			select {
			case cleaner := <-done:
				log.Println("Flushin to disk")
				cleaner.WaitGroup.Done()
				return

			case <-writeTic:
				d.flush()
			}
		}
	}()

	return done
}

func (d *AccessClient) flush() {
	d.DataStore(func(data *DataStore) {
		log.Printf("Flushing data to file: %s\n", d.dbname)
		file, err := os.Create(d.dbname)
		if err != nil {
			log.Println("Unable to write file to disk", err)
			return
		}
		defer file.Close()
		_, err = data.WriteTo(file)
		if err != nil {
			log.Println("Error occured while writing data", err)
		}
	})
}

func (d *AccessClient) DataStore(fn access) {
	defer d.lock.Unlock()
	d.lock.Lock()
	fn(d.store)
}

func (d *AccessClient) replaceItem(item Item) {
	d.DataStore(func(storage *DataStore) {
		for i, t := range storage.data.Items {
			if t.Id == item.Id {
				storage.data.Items[i] = item
			}
		}
	})
}

func (a *AccessClient) Users() (users []User) {
	a.DataStore(func(storage *DataStore) {
		users = make([]User, len(storage.data.Users))
		copy(users, storage.data.Users)
	})
	return
}

func (a *AccessClient) AddUser(u User) error {
	a.DataStore(func(storage *DataStore) {
		storage.data.Users = append(storage.data.Users, u)
	})
	return nil
}

func (a *AccessClient) UpdateUser(updated User) (user User, err error) {
	a.DataStore(func(storage *DataStore) {
		for i := 0; i < len(storage.data.Users); i++ {
			t := &storage.data.Users[i]
			if t.Id == updated.Id {
				t.Username = updated.Username
				t.Email = updated.Email
				t.FirstName = updated.FirstName
				t.LastName = updated.LastName
				user, err = *t, nil
				return
			}
		}
		user, err = User{}, errors.New("Couldn't find item to update")
	})
	return
}

func (a *AccessClient) FindUser(id string) (user User, err error) {
	a.DataStore(func(storage *DataStore) {
		for _, u := range storage.data.Users {
			if u.Id == id {
				user, err = u, nil
				return
			}
		}
		user, err = User{}, errors.New("Unable to find user")
	})
	return
}

func (a *AccessClient) FindItem(id string) (item Item, err error) {
	a.DataStore(func(storage *DataStore) {
		for _, t := range storage.data.Items {
			if t.Id == id {
				item, err = t, nil
				return
			}
		}
		item, err = Item{}, errors.New("Couln't find item")
	})
	return
}

func (a *AccessClient) UpdateItem(id string, updated Item) (item Item, err error) {
	a.DataStore(func(storage *DataStore) {
		for i := 0; i < len(storage.data.Items); i++ {
			t := &storage.data.Items[i]
			if t.Id == id {
				t.Title = updated.Title
				t.Summary = updated.Summary
				item, err = *t, nil
				return
			}
		}
		item, err = Item{}, errors.New("Couldn't find item to update")
	})
	return
}

func (a *AccessClient) AddItem(item Item) error {
	a.DataStore(func(storage *DataStore) {
		storage.data.Items = append(storage.data.Items, item)
	})
	return nil
}

func (d *AccessClient) Items() (res []Item) {
	d.DataStore(func(storage *DataStore) {
		res = make([]Item, len(storage.data.Items))
		copy(res, storage.data.Items)
	})
	return
}
