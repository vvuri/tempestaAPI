package localdb

import (
	"log"

	"github.com/boltdb/bolt"
)

type Storage struct {
	basePath string
}

func New(basePath string) Storage {
	return Storage{basePath: basePath}
}

func (c Storage) Save(page *storage.Page) (err error) {
	db, err := bolt.Open("telepages.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}
