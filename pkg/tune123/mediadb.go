package tune123

import (
	"errors"
	"fmt"
	"path/filepath"
)

type DataRec struct {
	ID         int64
	FolderName string
	FileName   string
	FullPath   string
}

type Database struct {
	Rec []DataRec
}

func (d *Database) BuildDatabase() error {
	d.Rec = d.Rec[:0] //clear database

	files, err := FindAllFiles(GLOBALCONFIG.AudioPath[0], GLOBALMASK)
	if err != nil {
		fmt.Println("BuildDatabase error:", err)
		return err
	}

	i := 0
	for k, v := range files {
		var rec DataRec
		rec.ID = int64(i)
		rec.FolderName = v //filepath.Dir(k)
		rec.FullPath = k
		rec.FileName = filepath.Base(k)
		fmt.Println("DB:", rec.ID, "-", rec.FolderName, "-", rec.FileName, "-", rec.FullPath)

		d.Rec = append(d.Rec, rec)
		i++
	}

	return err
}

func (d *Database) RecordExist(recordID int64) bool {
	for _, r := range d.Rec {
		if r.ID != recordID {
			continue
		}
		return true
	}
	return false
}

func (d *Database) Record(recordID int64) (record DataRec, err error) {
	if !d.RecordExist(recordID) {
		return record, errors.New("Запись отсутствует в аудиотеке")
	}

	for _, r := range d.Rec {
		if r.ID != recordID {
			continue
		}
		return r, nil
	}

	return record, err
}
