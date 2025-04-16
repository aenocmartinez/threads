package di

import (
	"database/sql"
	"sync"

	"threads/src/infraestructure/dao"
	"threads/src/infraestructure/database"
)

type Container struct {
	db       *sql.DB
	userRepo *dao.UserDAO
}

var (
	instance *Container
	once     sync.Once
)

func GetContainer() *Container {
	once.Do(func() {
		db := database.GetDB()
		instance = &Container{
			db:       db,
			userRepo: dao.NewUserDAO(db),
		}
	})
	return instance
}

func (c *Container) GetUserRepository() *dao.UserDAO {
	return c.userRepo
}
