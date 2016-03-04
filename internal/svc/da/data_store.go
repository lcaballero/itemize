package da

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

type DataStore struct {
	data Data
}

func LoadFromFile(dbname string) (*DataStore, error) {
	bits, err := ioutil.ReadFile(dbname)
	if err != nil {
		return nil, err
	}
	data := &Data{}
	err = json.Unmarshal(bits, data)
	if err != nil {
		return nil, err
	}
	d := &DataStore{
		data: data,
	}
	return d, nil
}

func NewDataStore() (*DataStore, error) {
	d := &DataStore{
		data: NewData().Add(NewUser()),
	}
	return d, nil
}

// WriteTo write the data to the given writer which can be loaded by
// LoadFromFile to recreate the DataStore from file.
func (d *DataStore) WriteTo(w io.Writer) (int64, error) {
	bits, err := json.Marshal(d.data)
	if err != nil {
		return 0, nil
	}
	return w.Write(bits)
}
