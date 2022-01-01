package persistence

import (
	"fmt"
	"food-app-go/domain/entity"
	"food-app-go/domain/repository"
	"github.com/jinzhu/gorm"
)

type Repositories struct {
	User repository.UserRepository

	db *gorm.DB
}

func NewRepositories(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*Repositories, error) {
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	db, err := gorm.Open(Dbdriver, DBURL)

	if err != nil {
		return nil, err
	}
	db.LogMode(true)

	return &Repositories{
		User: NewUserRepository(db),
		db:   db,
	}, nil
}

func (r *Repositories) Close() error {
	return r.db.Close()
}

func (r *Repositories) Automigrate() error {
	return r.db.AutoMigrate(&entity.User{}).Error
}
