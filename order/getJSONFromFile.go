package order

import (
	"encoding/json"
	"errors"
)

func (db *FileDataBase) GetJSONFromFile() ([]Order, error) {

	bytesFile, err := getBytesFromFile(fileName)
	if err != nil {
		return db.data, errors.New(statusServerError)
	}

	err = json.Unmarshal(bytesFile, &db.data)
	if err != nil {
		return db.data, errors.New(statusServerError)
	}
	return db.data, nil
}
