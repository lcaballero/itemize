package da

import (
	"log"
	"os"
	"time"

	"github.com/lcaballero/hitman"
)

type DataWriter struct {
	onActiveStore chan *DataStore
	data          *DataStore
	dbname        string
}

func NewDataWriter(dbname string, activated chan *DataStore) (*DataWriter, error) {
	d := &DataWriter{
		onActiveStore: activated,
		data:          nil,
		dbname:        dbname,
	}
	return d, nil
}

func (d *DataWriter) Start() hitman.KillChannel {
	done := hitman.NewKillChannel()
	writeTic := time.NewTicker(5 * time.Second).C

	go func() {
		d.onActiveStore <- d.data
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

func (d *DataWriter) flush() {
	file, err := os.Create(d.dbname)
	if err != nil {
		log.Println("Unable to write file to disk", err)
		return
	}
	defer file.Close()
	_, err = d.data.WriteTo(file)
	if err != nil {
		log.Println("Error occured while writing data", err)
	}
}
