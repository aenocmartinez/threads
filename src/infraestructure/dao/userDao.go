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
		SELECT id, name, username, email,
			COALESCE(phone, '0'), password,
			COALESCE(avatar, ''), COALESCE(description, ''),
			COALESCE(session_token, '')
		FROM users
		WHERE id = $1`

	row := u.db.QueryRow(query, id)

	var userID int64
	var name, username, email, phone, password, avatar, description, sessionToken string

	err := row.Scan(&userID, &name, &username, &email, &phone, &password, &avatar, &description, &sessionToken)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.NewUser(u), nil
		}
		return nil, fmt.Errorf("error consultando usuario: %w", err)
	}

	user := domain.NewUser(u)
	user.SetID(userID)
	user.SetName(name)
	user.SetUsername(username)
	user.SetEmail(email)
	user.SetPhone(phone)
	user.SetPassword(password)
	user.SetAvatar(avatar)
	user.SetDescription(description)
	user.SetSessionToken(sessionToken)

	return user, nil
}

func (u *UserDAO) FindByEmail(email string) (*domain.User, error) {
	query := `
		SELECT id, name, username, email,
			COALESCE(phone, '0'), password,
			COALESCE(avatar, ''), COALESCE(description, ''),
			COALESCE(session_token, '')
		FROM users
		WHERE email = $1`

	row := u.db.QueryRow(query, email)

	var id int64
	var name, username, phone, password, avatar, description, sessionToken string

	err := row.Scan(&id, &name, &username, &email, &phone, &password, &avatar, &description, &sessionToken)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.NewUser(u), nil
		}
		return nil, fmt.Errorf("error buscando usuario por email: %w", err)
	}

	user := domain.NewUser(u)
	user.SetID(id)
	user.SetName(name)
	user.SetUsername(username)
	user.SetEmail(email)
	user.SetPhone(phone)
	user.SetPassword(password)
	user.SetAvatar(avatar)
	user.SetDescription(description)
	user.SetSessionToken(sessionToken)

	return user, nil
}

func (u *UserDAO) FindUserLogin(login string) (*domain.User, error) {
	query := `
		SELECT id, name, username, email,
			COALESCE(phone, '0'), password,
			COALESCE(avatar, ''), COALESCE(description, ''),
			COALESCE(session_token, '')
		FROM users
		WHERE username = $1 OR email = $2 OR phone = $3`

	row := u.db.QueryRow(query, login, login, login)

	var id int64
	var name, username, email, phone, password, avatar, description, sessionToken string

	err := row.Scan(&id, &name, &username, &email, &phone, &password, &avatar, &description, &sessionToken)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.NewUser(u), nil
		}
		return nil, fmt.Errorf("error buscando usuario por login: %w", err)
	}

	user := domain.NewUser(u)
	user.SetID(id)
	user.SetName(name)
	user.SetUsername(username)
	user.SetEmail(email)
	user.SetPhone(phone)
	user.SetPassword(password)
	user.SetAvatar(avatar)
	user.SetDescription(description)
	user.SetSessionToken(sessionToken)

	return user, nil
}

func (u *UserDAO) FindByUsername(username string) (*domain.User, error) {
	query := `
		SELECT id, name, username, email,
			COALESCE(phone, '0'), password,
			COALESCE(avatar, ''), COALESCE(description, ''),
			COALESCE(session_token, '')
		FROM users
		WHERE username = $1`

	row := u.db.QueryRow(query, username)

	var id int64
	var name, email, phone, password, avatar, description, sessionToken string

	err := row.Scan(&id, &name, &username, &email, &phone, &password, &avatar, &description, &sessionToken)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.NewUser(u), nil
		}
		return nil, fmt.Errorf("error buscando usuario por username: %w", err)
	}

	user := domain.NewUser(u)
	user.SetID(id)
	user.SetName(name)
	user.SetUsername(username)
	user.SetEmail(email)
	user.SetPhone(phone)
	user.SetPassword(password)
	user.SetAvatar(avatar)
	user.SetDescription(description)
	user.SetSessionToken(sessionToken)

	return user, nil
}

func (u *UserDAO) Save(user *domain.User) error {
	query := `
		INSERT INTO users (name, username, email, phone, password, avatar, description, session_token)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id`

	// Preparar valores NULL para campos vacíos
	var phone, avatar, description, sessionToken interface{}

	if user.GetPhone() == "" {
		phone = nil
	} else {
		phone = user.GetPhone()
	}

	if user.GetAvatar() == "" {
		avatar = nil
	} else {
		avatar = user.GetAvatar()
	}

	if user.GetDescription() == "" {
		description = nil
	} else {
		description = user.GetDescription()
	}

	if user.GetSessionToken() == "" {
		sessionToken = nil
	} else {
		sessionToken = user.GetSessionToken()
	}

	var insertedID int64
	err := u.db.QueryRow(query,
		user.GetName(),
		user.GetUsername(),
		user.GetEmail(),
		phone,
		user.GetPassword(),
		avatar,
		description,
		sessionToken,
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
		UPDATE users SET 
			name = $1,
			username = $2,
			email = $3,
			phone = $4,
			password = $5,
			avatar = $6,
			description = $7,
			session_token = $8
		WHERE id = $9`

	// Preparar valores NULL para campos vacíos
	var phone, avatar, description, sessionToken interface{}

	if user.GetPhone() == "" {
		phone = nil
	} else {
		phone = user.GetPhone()
	}

	if user.GetAvatar() == "" {
		avatar = nil
	} else {
		avatar = user.GetAvatar()
	}

	if user.GetDescription() == "" {
		description = nil
	} else {
		description = user.GetDescription()
	}

	if user.GetSessionToken() == "" {
		sessionToken = nil
	} else {
		sessionToken = user.GetSessionToken()
	}

	_, err := u.db.Exec(query,
		user.GetName(),
		user.GetUsername(),
		user.GetEmail(),
		phone,
		user.GetPassword(),
		avatar,
		description,
		sessionToken,
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

func (u *UserDAO) ExistsUsername(username string) (bool, error) {
	query := `SELECT 1 FROM users WHERE username = $1 LIMIT 1`
	row := u.db.QueryRow(query, username)

	var dummy int
	err := row.Scan(&dummy)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("error verificando existencia de username: %w", err)
	}
	return true, nil
}
