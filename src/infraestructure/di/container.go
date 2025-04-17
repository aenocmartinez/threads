package di

import (
	"database/sql"
	"sync"

	"threads/src/domain"
	"threads/src/infraestructure/dao"
	"threads/src/infraestructure/database"
)

type Container struct {
	db             *sql.DB
	userRepo       domain.UserRepository
	comentarioRepo domain.ComentarioRepository
}

var (
	instance *Container
	once     sync.Once
)

func GetContainer() *Container {
	once.Do(func() {
		db := database.GetDB()
		instance = &Container{
			db:             db,
			userRepo:       dao.NewUserDAO(db),
			comentarioRepo: dao.NewComentarioDAO(db),
		}
	})
	return instance
}

func (c *Container) GetUserRepository() domain.UserRepository {
	return c.userRepo
}

func (c *Container) GetComentarioRepository() domain.ComentarioRepository {
	return c.comentarioRepo
}
