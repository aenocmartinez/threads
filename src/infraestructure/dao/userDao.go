package dao

import (
	"database/sql"
	"fmt"
	"threads/src/domain"
)

type UserDAO struct {
	db *sql.DB
}

func NewUserDAO(db *sql.DB) *UserDAO {
	return &UserDAO{db: db}
}

func (u *UserDAO) FindByID(id int64) (*domain.User, error) {
	query := `
		SELECT id, username, email, COALESCE(phone, '0'), password, COALESCE(session_token, '')
		FROM users
		WHERE id = $1
	`

	row := u.db.QueryRow(query, id)

	var userID int64
	var username, email, phone, password, sessionToken string

	err := row.Scan(&userID, &username, &email, &phone, &password, &sessionToken)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.NewUser(u), nil
		}
		return nil, fmt.Errorf("error consultando usuario: %w", err)
	}

	user := domain.NewUser(u)
	user.SetID(userID)
	user.SetUsername(username)
	user.SetEmail(email)
	user.SetPhone(phone)
	user.SetPassword(password)
	user.SetSessionToken(sessionToken)

	return user, nil
}

func (u *UserDAO) FindByEmail(email string) (*domain.User, error) {
	query := `
		SELECT id, username, email, COALESCE(phone, '0'), password, COALESCE(session_token, '')
		FROM users
		WHERE email = $1
	`

	row := u.db.QueryRow(query, email)

	var id int64
	var username, phone, password, sessionToken string

	err := row.Scan(&id, &username, &email, &phone, &password, &sessionToken)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.NewUser(u), nil
		}
		return nil, fmt.Errorf("error buscando usuario por email: %w", err)
	}

	user := domain.NewUser(u)
	user.SetID(id)
	user.SetUsername(username)
	user.SetEmail(email)
	user.SetPhone(phone)
	user.SetPassword(password)
	user.SetSessionToken(sessionToken)

	return user, nil
}

func (u *UserDAO) FindUserLogin(login string) (*domain.User, error) {
	query := `
		SELECT id, username, email, COALESCE(phone, '0'), password, COALESCE(session_token, '')
		FROM users
		WHERE username = $1 OR email = $2 OR phone = $3
	`

	row := u.db.QueryRow(query, login, login, login)

	var id int64
	var username, email, phone, password, sessionToken string

	err := row.Scan(&id, &username, &email, &phone, &password, &sessionToken)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.NewUser(u), nil
		}
		return nil, fmt.Errorf("error buscando usuario por login: %w", err)
	}

	user := domain.NewUser(u)
	user.SetID(id)
	user.SetUsername(username)
	user.SetEmail(email)
	user.SetPhone(phone)
	user.SetPassword(password)
	user.SetSessionToken(sessionToken)

	return user, nil
}

func (u *UserDAO) FindByUsername(username string) (*domain.User, error) {
	query := `
		SELECT id, username, email, COALESCE(phone, '0'), password, COALESCE(session_token, '')
		FROM users
		WHERE username = $1
	`

	row := u.db.QueryRow(query, username)

	var id int64
	var email, phone, password, sessionToken string

	err := row.Scan(&id, &username, &email, &phone, &password, &sessionToken)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.NewUser(u), nil
		}
		return nil, fmt.Errorf("error buscando usuario por username: %w", err)
	}

	user := domain.NewUser(u)
	user.SetID(id)
	user.SetUsername(username)
	user.SetEmail(email)
	user.SetPhone(phone)
	user.SetPassword(password)
	user.SetSessionToken(sessionToken)

	return user, nil
}

func (u *UserDAO) Save(user *domain.User) error {
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id`

	var insertedID int64
	err := u.db.QueryRow(query,
		user.GetUsername(),
		user.GetEmail(),
		user.GetPassword(),
	).Scan(&insertedID)

	if err != nil {
		return fmt.Errorf("error al guardar usuario: %w", err)
	}

	user.SetID(insertedID)
	return nil
}

func (u *UserDAO) Update(user *domain.User) error {
	if user.GetID() <= 0 {
		return fmt.Errorf("el ID del usuario no es válido")
	}

	query := `
		UPDATE users 
		SET username = $1, email = $2, phone = $3, password = $4, session_token = $5 
		WHERE id = $6
	`

	_, err := u.db.Exec(query,
		user.GetUsername(),
		user.GetEmail(),
		user.GetPhone(),
		user.GetPassword(),
		user.GetSessionToken(),
		user.GetID(),
	)
	if err != nil {
		return fmt.Errorf("error al actualizar usuario con ID %d: %w", user.GetID(), err)
	}

	return nil
}

func (u *UserDAO) Delete(id int64) error {
	if id <= 0 {
		return fmt.Errorf("el ID del usuario no es válido")
	}

	query := `DELETE FROM users WHERE id = $1`

	_, err := u.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar usuario con ID %d: %w", id, err)
	}

	return nil
}
