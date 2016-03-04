package da

import (
	"log"
	"os"
	"time"

	"github.com/lcaballero/hitman"
)

const DefaultDbName = "items.db.json"

type DataWriter struct {
	onActiveStore chan *DataStore
	data          *DataStore
	dbname        string
}

func NewDataWriter(dbname string, activated chan *DataStore) (*DataWriter, error) {
	// TODO: check if DB file exists so that dbname is the 'intended' name.
	data, err := LoadDataStore(dbname)
	if err != nil {
		return nil, err
	}

	if dbname == "" {
		dbname = DefaultDbName
	}
	d := &DataWriter{
		onActiveStore: activated,
		data:          data,
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
	log.Printf("Flushing data to file: %s\n", d.dbname)
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
