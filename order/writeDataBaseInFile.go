package order

import (
	"encoding/json"
)

func (db *FileDataBase) writeDataBaseInFile() error {
	bytesOrder, err := json.Marshal(db.data)
	if err != nil {
		return err
	}

	err = writeTextInFile(fileName, bytesOrder)
	if err != nil {
		return err
	}
	return nil
}
