package mysql

import (
	"database/sql"
	"fmt"
	"threads/src/domain"
	"time"
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
		WHERE id = ?`

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
		WHERE email = ?`

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
		FROM users WHERE username = ? OR email = ? OR phone = ?`

	fmt.Println(query)
	fmt.Println(login)

	row := u.db.QueryRow(query, login, login, login)

	var id int64
	var name, username, email, phone, password, avatar, description, sessionToken string

	err := row.Scan(&id, &name, &username, &email, &phone, &password, &avatar, &description, &sessionToken)
	if err != nil {
		fmt.Println("err: ", err)
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
		WHERE username = ?`

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
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

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

	result, err := u.db.Exec(query,
		user.GetName(),
		user.GetUsername(),
		user.GetEmail(),
		phone,
		user.GetPassword(),
		avatar,
		description,
		sessionToken,
	)
	if err != nil {
		return fmt.Errorf("error al guardar usuario: %w", err)
	}

	insertedID, err := result.LastInsertId()
	if err == nil {
		user.SetID(insertedID)
	}

	return nil
}

func (u *UserDAO) Update(user *domain.User) error {
	if user.GetID() <= 0 {
		return fmt.Errorf("el ID del usuario no es válido")
	}

	query := `
		UPDATE users SET 
			name = ?,
			username = ?,
			email = ?,
			phone = ?,
			password = ?,
			avatar = ?,
			description = ?,
			session_token = ?
		WHERE id = ?`

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

	query := `DELETE FROM users WHERE id = ?`

	_, err := u.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar usuario con ID %d: %w", id, err)
	}

	return nil
}

func (u *UserDAO) ExistsUsername(username string) (bool, error) {
	query := `SELECT 1 FROM users WHERE username = ? LIMIT 1`
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

func (u *UserDAO) ObtenerUsuariosQueMeSiguen(userID int64) []domain.Seguidor {
	query := `
		SELECT 
			us.id, us.name, us.username, us.email, 
			COALESCE(us.phone, ''),
			us.password,
			COALESCE(us.avatar, ''),
			COALESCE(us.description, ''),
			COALESCE(us.session_token, ''),
			s.fecha_sigue
		FROM seguidores s
		JOIN users us ON us.id = s.usuario_seguidor_id
		WHERE s.usuario_seguido_id = ?
	`

	rows, err := u.db.Query(query, userID)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer rows.Close()

	seguidores := []domain.Seguidor{}

	for rows.Next() {
		var (
			id           int64
			name         string
			username     string
			email        string
			phone        string
			password     string
			avatar       string
			description  string
			sessionToken string
			fechaSigue   time.Time
		)

		err := rows.Scan(&id, &name, &username, &email, &phone, &password, &avatar, &description, &sessionToken, &fechaSigue)
		if err != nil {
			fmt.Println(err)
			continue
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

		seguidor := domain.NewSeguidor()
		seguidor.SetUserSeguidor(user)
		seguidor.SetFechaSigue(fechaSigue)

		seguidores = append(seguidores, *seguidor)
	}

	return seguidores
}

func (u *UserDAO) ObtenerUsuariosQueSigo(userID int64) []domain.Seguidor {
	query := `
		SELECT 
			us.id,
			us.name,
			us.username,
			us.email,
			COALESCE(us.phone, ''),
			us.password,
			COALESCE(us.avatar, ''),
			COALESCE(us.description, ''),
			COALESCE(us.session_token, ''),
			s.fecha_sigue
		FROM seguidores s
		JOIN users us ON us.id = s.usuario_seguido_id
		WHERE s.usuario_seguidor_id = ?
	`

	rows, err := u.db.Query(query, userID)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer rows.Close()

	seguidos := []domain.Seguidor{}

	for rows.Next() {
		var (
			id           int64
			name         string
			username     string
			email        string
			phone        string
			password     string
			avatar       string
			description  string
			sessionToken string
			fechaSigue   time.Time
		)

		err := rows.Scan(&id, &name, &username, &email, &phone, &password, &avatar, &description, &sessionToken, &fechaSigue)
		if err != nil {
			fmt.Println(err)
			continue
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

		seguidor := domain.NewSeguidor()
		seguidor.SetUserSeguido(user)
		seguidor.SetFechaSigue(fechaSigue)

		seguidos = append(seguidos, *seguidor)
	}

	return seguidos
}

func (u *UserDAO) SeguirUsuario(usuarioSeguidorID, usuarioSeguidoID int64) bool {
	query := `
		INSERT IGNORE INTO seguidores (usuario_seguidor_id, usuario_seguido_id)
		VALUES (?, ?)`

	_, err := u.db.Exec(query, usuarioSeguidorID, usuarioSeguidoID)
	return err == nil
}

func (u *UserDAO) DejarDeSeguirUsuario(usuarioSeguidorID, usuarioSeguidoID int64) bool {
	query := `
		DELETE FROM seguidores
		WHERE usuario_seguidor_id = ? AND usuario_seguido_id = ?`

	_, err := u.db.Exec(query, usuarioSeguidorID, usuarioSeguidoID)
	return err == nil
}

func (u *UserDAO) TotalNumeroDeSeguidores(usuarioID int64) int {
	query := `
		SELECT COUNT(*)
		FROM seguidores
		WHERE usuario_seguido_id = ?`

	var total int
	err := u.db.QueryRow(query, usuarioID).Scan(&total)
	if err != nil {
		fmt.Println("error contando seguidores:", err)
		return 0
	}

	return total
}
