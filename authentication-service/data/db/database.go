package db

import (
	"fmt"

	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Database struct {
	DB *gorm.DB
}

func NewDB() (Database, error) {
	//dsn := "bac4178dc89368:292965a5@tcp(us-cdbr-east-05.cleardb.net)/heroku_560fb6556eff9f8?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "host=localhost port=5432 user=postgres password=password dbname=users sslmode=disable timezone=UTC connect_timeout=5"

	db, err := gorm.Open("postgres", dsn)
	return Database{
		DB: db,
	}, err
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

func (db *Database) FindWithPagination(condition interface{}, offset int, pageSize int, value interface{}) error {
	err := db.DB.Offset(offset).Limit(pageSize).Find(value, condition).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil
	}
	return err
}

func (db *Database) Create(value interface{}) error {
	err := db.DB.Create(value).Error
	return err
}

func (db *Database) CreateMany(value interface{}, model interface{}) error {
	fmt.Println(model)
	err := db.DB.Model(model).Create(value).Error
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
