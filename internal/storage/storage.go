package storage

type Storage struct {
	data map[int]string
}

func NewStorage() *Storage {
	storage := Storage{
		data: make(map[int]string),
	}
	storage.data[1] = "You are perfect because of your imperfections."
	storage.data[2] = "Do what inspires you. Life is too short not to love the job you do every day."
	storage.data[3] = "Complaining will not get anything done."

	return &storage
}

func (s Storage) GetRandom() (st string) {

	for _, st = range s.data {
		return st
	}

	return ""
}
