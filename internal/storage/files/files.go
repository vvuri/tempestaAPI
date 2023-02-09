package files

import (
	"encoding/gob"
	"os"
	"path/filepath"
	"tempestaAPI/internal/storage"
)

const defaultFilePerm = 0774

type Storage struct {
	basePath string
}

func New(basePath string) Storage {
	return Storage{basePath: basePath}
}

func (s Storage) Save(page *storage.Page) (err error) {
	defer func() { err = e.WrapIfErr("can't save page") }()

	fPath := filepath.Join(s.basePath, page.UserName)

	if err := os.MkdirAll(filePath, defaultFilePerm); err != nil {
		return err
	}

	fName, err := fileName(page)
	if err != nil {
		return err
	}

	fPath = filepath.Join(fPath, fName)

	file, err := os.Create(fPath)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := gob.NewEncoder(file).Encode(page); err != nil {
		return err
	}

	return nil
}

func fileName(p *storage.Page) (string, error) {
	return p.Hash(), nil
}
