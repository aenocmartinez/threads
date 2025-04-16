package domain

type UserRepository interface {
	FindByID(id int64) (*User, error)
	FindByEmail(email string) (*User, error)
	FindByUsername(username string) (*User, error)
	FindUserLogin(login string) (*User, error)
	ExistsUsername(username string) (bool, error)
	Save(user *User) error
	Update(user *User) error
	Delete(id int64) error
}
