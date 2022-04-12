package datastore

import (
	"fmt"
	"gormbe8/entities"

	"gorm.io/gorm"
)

type UserDB struct {
	Db *gorm.DB
}

func (u *UserDB) GetAllDataUser() ([]entities.User, error) {
	res := []entities.User{}

	if err := u.Db.Table("user").Where("Nama LIKE ?", "%j%").Find(&res).Error; err != nil {
		fmt.Println("Terjadi kesalahan saat get data user", err)
		return []entities.User{}, err
	}

	return res, nil
}
