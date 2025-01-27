package order

func (db *InMemoryDataBase) GetStatus() (string, error) {
	status := "ok"
	return status, nil
}
