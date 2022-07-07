package db

import (
	"authentication/data/entity"

	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/jinzhu/gorm"
	// import mysql driver
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	DB *gorm.DB
}

func NewDB() (Database, error) {
	//dsn := "bac4178dc89368:292965a5@tcp(us-cdbr-east-05.cleardb.net)/heroku_560fb6556eff9f8?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "hiendv:password@tcp(mysql:3306)/usersdb?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	return Database{
		DB: db,
	}, err
}
func (db *Database) Migrate() error {
	err := db.DB.AutoMigrate(entity.User{}).Error

	return err
}
func (db *Database) First(condition interface{}, value interface{}) error {
	err := db.DB.First(value, condition).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil
	}

	return err
}
func (db *Database) Find(condition interface{}, value interface{}) error {
	err := db.DB.Find(value, condition).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil
	}
	return err
}

func (db *Database) Create(value interface{}) error {
	err := db.DB.Create(value).Error
	return err
}

func (db *Database) Delete(value interface{}) error {
	return db.DB.Delete(value).Error
}

func (db *Database) Update(model interface{}, oldVal interface{}, newVal interface{}) error {
	return db.DB.Model(model).Where(oldVal).Updates(newVal).Error
}

func (db *Database) ExcQuery(query string) error {
	return db.DB.Exec(query).Error
}
