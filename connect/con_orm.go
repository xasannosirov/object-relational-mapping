package connect

import (
	mod "gorm/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// This function accepts user, password, dbname and connects to GORM,
// if it can connect it returns *gorm.DB and nil otherwise *gorm.DB and error
// default values host=localhost, TimeZone=Asia/Tashkent, sslmode=disable
func ConnectToGORM(user, password, dbname string) (*gorm.DB, error) {
	dsn := "host=localhost " +
		"user=" + user + " " +
		"password=" + password + " " +
		"dbname=" + dbname + " " +
		"sslmode=disable " +
		"TimeZone=Asia/Tashkent"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// gorm connect eror checking
	if err != nil {
		return &gorm.DB{}, err
	}

	return db, nil
}

// this func auto migrate 
func AutoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&mod.Students{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&mod.Courses{}); err != nil {
		return err
	}
	return nil
}