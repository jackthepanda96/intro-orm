package main

import (
	"fmt"
	"gormbe8/datastore"
	"gormbe8/entities"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Username string
	Password string
	Host     string
	Port     int16
	DB       string
}

func ConnectDB(configData Config) *gorm.DB {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configData.Username,
		configData.Password,
		configData.Host,
		configData.Port,
		configData.DB)
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		fmt.Println("Terjadi kesalahan saat koneksi database", err)
		return nil
	}
	return db
}

func ReadEnv() Config {
	if err := godotenv.Load("local.env"); err != nil {
		fmt.Println("ERROR LOAD FILE", err)
	}

	res := Config{}
	res.Username = os.Getenv("User")
	res.DB = os.Getenv("DB")
	res.Password = os.Getenv("Password")
	res.Host = os.Getenv("Host")
	intConv, _ := strconv.Atoi(os.Getenv("Port"))
	res.Port = int16(intConv)
	return res
}

func main() {

	config := ReadEnv()

	// Koneksi ke database
	db := ConnectDB(config)
	// // db.AutoMigrate(&Barang{})

	fmt.Println(db)
	fmt.Println(db.Error)

	userAcc := datastore.UserDB{Db: db}
	barangAcc := datastore.BarangDB{Db: db}

	allUser, err := userAcc.GetAllDataUser()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(allUser)

	_, err = barangAcc.InsertBarang(entities.Barang{Nama: "Blender Kithcen Aid"})

	allBarang, err := barangAcc.GetAllDataBarang()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(allBarang)

}
