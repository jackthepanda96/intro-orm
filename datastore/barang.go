package datastore

import (
	"fmt"
	"gormbe8/entities"

	"gorm.io/gorm"
)

type BarangDB struct {
	Db *gorm.DB
}

func (b *BarangDB) GetAllDataBarang() ([]entities.Barang, error) {
	res := []entities.Barang{}

	if err := b.Db.Find(&res).Error; err != nil {
		fmt.Println("Terjadi kesalahan saat get data barang", err)
		return []entities.Barang{}, err
	}

	return res, nil
}

func (b *BarangDB) InsertBarang(newBarang entities.Barang) (entities.Barang, error) {
	if err := b.Db.Create(&newBarang).Error; err != nil {
		fmt.Println("Terjadi kesalahan saat insert barang", err)
		return newBarang, err
	}
	return newBarang, nil
}
