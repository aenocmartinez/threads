package dao

import (
	"database/sql"
	"fmt"
	"threads/src/domain"
	"threads/src/view/dto"
	"time"
)

type ComentarioDAO struct {
	db *sql.DB
}

func NewComentarioDAO(db *sql.DB) *ComentarioDAO {
	return &ComentarioDAO{db: db}
}

func (c *ComentarioDAO) CrearComentario(comentario *domain.Comentario) bool {
	query := `
		INSERT INTO comentarios (usuario_id, contenido, comentario_padre_id, created_at)
		VALUES ($1, $2, $3, CURRENT_TIMESTAMP)
		RETURNING id
	`

	var id int64
	var padreID *int64

	if comentario.GetComentarioPadre() != nil {
		tmpID := comentario.GetComentarioPadre().GetID()
		padreID = &tmpID
	}

	err := c.db.QueryRow(
		query,
		comentario.GetUser().GetID(),
		comentario.GetContenido(),
		padreID,
	).Scan(&id)

	if err != nil {
		fmt.Println("error creando comentario:", err)
		return false
	}

	comentario.SetID(id)
	return true
}

func (c *ComentarioDAO) ResponderAComentario(comentarioPadreID int64, respuesta *domain.Comentario) bool {
	query := `
		INSERT INTO comentarios (usuario_id, contenido, comentario_padre_id, created_at)
		VALUES ($1, $2, $3, CURRENT_TIMESTAMP)
		RETURNING id
	`

	var id int64

	err := c.db.QueryRow(
		query,
		respuesta.GetUser().GetID(),
		respuesta.GetContenido(),
		comentarioPadreID,
	).Scan(&id)

	if err != nil {
		fmt.Println("error respondiendo comentario:", err)
		return false
	}

	respuesta.SetID(id)
	return true
}

func (c *ComentarioDAO) ActualizarComentario(comentario *domain.Comentario) bool {
	query := `
		UPDATE comentarios
		SET contenido = $1,
		    updated_at = CURRENT_TIMESTAMP
		WHERE id = $2
	`

	_, err := c.db.Exec(query, comentario.GetContenido(), comentario.GetID())
	if err != nil {
		fmt.Println("error actualizando comentario:", err)
		return false
	}

	return true
}

func (c *ComentarioDAO) EliminarComentario(comentarioID int64) bool {
	query := `DELETE FROM comentarios WHERE id = $1`

	_, err := c.db.Exec(query, comentarioID)
	if err != nil {
		fmt.Println("error eliminando comentario:", err)
		return false
	}

	return true
}

func (c *ComentarioDAO) ObtenerComentario(id int64) *domain.Comentario {
	query := `
		SELECT usuario_id, contenido, comentario_padre_id, created_at, updated_at
		FROM comentarios
		WHERE id = $1
	`

	row := c.db.QueryRow(query, id)

	var (
		usuarioID         int64
		contenido         string
		comentarioPadreID *int64
		createdAt         time.Time
		updatedAt         *time.Time
	)

	err := row.Scan(&usuarioID, &contenido, &comentarioPadreID, &createdAt, &updatedAt)
	if err == sql.ErrNoRows {
		return domain.NewComentario(NewUserDAO(c.db), c)
	}
	if err != nil {
		fmt.Println("error consultando comentario:", err)
		return domain.NewComentario(NewUserDAO(c.db), c)
	}

	comentario := domain.NewComentario(NewUserDAO(c.db), c)
	comentario.SetID(id)
	comentario.SetContenido(contenido)
	comentario.SetCreatedAt(createdAt)
	if updatedAt != nil {
		comentario.SetUpdatedAt(*updatedAt)
	}

	user, _ := NewUserDAO(c.db).FindByID(usuarioID)
	if user != nil && user.Exists() {
		comentario.SetUser(user)
	}

	if comentarioPadreID != nil {
		padre := c.ObtenerComentario(*comentarioPadreID)
		if padre.Existe() {
			comentario.SetComentarioPadre(padre)
		}
	}

	return comentario
}

func (c *ComentarioDAO) ObtenerConversacion(comentarioID int64) dto.ComentarioConRespuestasDTO {
	comentario := c.ObtenerComentario(comentarioID)

	dtoConversacion := dto.ComentarioConRespuestasDTO{
		Comentario: *comentario.ToDTO(),
		Respuestas: c.obtenerRespuestasRecursivas(comentario.GetID()),
	}

	return dtoConversacion
}

func (c *ComentarioDAO) obtenerRespuestasRecursivas(comentarioID int64) []*dto.ComentarioConRespuestasDTO {
	query := `
		SELECT id
		FROM comentarios
		WHERE comentario_padre_id = $1
		ORDER BY created_at ASC
	`

	rows, err := c.db.Query(query, comentarioID)
	if err != nil {
		fmt.Println("error consultando respuestas del comentario:", err)
		return nil
	}
	defer rows.Close()

	respuestas := []*dto.ComentarioConRespuestasDTO{}

	for rows.Next() {
		var respuestaID int64
		if err := rows.Scan(&respuestaID); err != nil {
			fmt.Println("error leyendo id de respuesta:", err)
			continue
		}

		respuesta := c.ObtenerComentario(respuestaID)
		item := &dto.ComentarioConRespuestasDTO{
			Comentario: *respuesta.ToDTO(),
			Respuestas: c.obtenerRespuestasRecursivas(respuestaID),
		}
		respuestas = append(respuestas, item)
	}

	return respuestas
}

func (c *ComentarioDAO) ObtenerComentariosRecientes() []dto.ComentarioConRespuestasDTO {
	query := `
		SELECT id
		FROM comentarios
		WHERE comentario_padre_id IS NULL
		ORDER BY created_at DESC
		LIMIT 50
	`

	rows, err := c.db.Query(query)
	if err != nil {
		fmt.Println("error consultando comentarios recientes:", err)
		return nil
	}
	defer rows.Close()

	conversaciones := []dto.ComentarioConRespuestasDTO{}

	for rows.Next() {
		var comentarioID int64
		if err := rows.Scan(&comentarioID); err != nil {
			fmt.Println("error leyendo id del comentario raíz:", err)
			continue
		}

		conversacion := c.ObtenerConversacion(comentarioID)
		conversaciones = append(conversaciones, conversacion)
	}

	return conversaciones
}

func (c *ComentarioDAO) ObtenerComentariosRecientesDesde(fechaUltimo time.Time) []dto.ComentarioConRespuestasDTO {
	var query string
	var rows *sql.Rows
	var err error

	if fechaUltimo.IsZero() {
		query = `
			SELECT id FROM comentarios
			WHERE comentario_padre_id IS NULL
			ORDER BY created_at DESC
			LIMIT 50
		`
		rows, err = c.db.Query(query)
	} else {
		query = `
			SELECT id FROM comentarios
			WHERE comentario_padre_id IS NULL AND created_at < $1
			ORDER BY created_at DESC
			LIMIT 50
		`
		rows, err = c.db.Query(query, fechaUltimo)
	}

	if err != nil {
		fmt.Println("error consultando comentarios recientes:", err)
		return nil
	}
	defer rows.Close()

	conversaciones := []dto.ComentarioConRespuestasDTO{}

	for rows.Next() {
		var comentarioID int64
		if err := rows.Scan(&comentarioID); err != nil {
			fmt.Println("error leyendo comentario_id:", err)
			continue
		}
		conversacion := c.ObtenerConversacion(comentarioID)
		conversaciones = append(conversaciones, conversacion)
	}

	return conversaciones
}

func (c *ComentarioDAO) DarMeGustaAComentario(usuarioID, comentarioID int64) bool {
	query := `
		INSERT INTO me_gusta_comentario (usuario_id, comentario_id)
		VALUES ($1, $2)
		ON CONFLICT DO NOTHING
	`
	_, err := c.db.Exec(query, usuarioID, comentarioID)
	if err != nil {
		fmt.Println("error insertando me gusta:", err)
		return false
	}
	return true
}

func (c *ComentarioDAO) QuitarMeGustaAComentario(usuarioID, comentarioID int64) bool {
	query := `
		DELETE FROM me_gusta_comentario
		WHERE usuario_id = $1 AND comentario_id = $2
	`
	_, err := c.db.Exec(query, usuarioID, comentarioID)
	if err != nil {
		fmt.Println("error quitando me gusta:", err)
		return false
	}
	return true
}
