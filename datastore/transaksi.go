package datastore

import (
	"fmt"
	"gormbe8/entities"

	"gorm.io/gorm"
)

type TransaksiDB struct {
	Db *gorm.DB
}

func (t *TransaksiDB) GetAllDataUser() ([]entities.Transaksi, error) {
	res := []entities.Transaksi{}

	if err := t.Db.Find(&res).Error; err != nil {
		fmt.Println("Terjadi kesalahan saat get data user", err)
		return []entities.Transaksi{}, err
	}

	return res, nil
}
