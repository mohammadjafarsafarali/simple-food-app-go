package persistence

import (
	"food-app-go/domain/entity"
	"food-app-go/domain/repository"
	"github.com/jinzhu/gorm"
)

type Repositories struct {
	User repository.UserRepository

	db *gorm.DB
}

func NewRepositories(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*Repositories, error) {
	DBURL := DbUser + ":" + DbPassword + "@" + "tcp(" + DbHost + ":" + DbPort + ")" + "/" + DbName + "?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true"

	db, err := gorm.Open("mysql", DBURL)

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
