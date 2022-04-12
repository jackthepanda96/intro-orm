package entities

import "gorm.io/gorm"

type Transaksi struct {
	gorm.Model
	NoNota           string
	AlamatPengiriman string
}
