package comentarios

import (
	"context"
	"database/sql"

	repo "project_user-management_ms/data"
	models "project_user-management_ms/models"
)

func NewSQLComentario(Conn *sql.DB) repo.Comentarios {
	return &mysqlComentarios{
		Conn: Conn,
	}
}

type mysqlComentarios struct {
	Conn *sql.DB
}

func (m *mysqlComentarios) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Comentarios, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.Comentarios, 0)
	for rows.Next() {
		data := new(models.Comentarios)

		err := rows.Scan(
			&data.IDComento,
			&data.IDComentado,
			&data.Comentario,
			&data.Fecha,
			&data.Hora,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (m *mysqlComentarios) Fetch(ctx context.Context, num int64) ([]*models.Comentarios, error) {
	query := "Select IDComento, IDComentado, Comentario, Fecha, Hora  From Comentarios limit ?"

	return m.fetch(ctx, query, num)
}

func (m *mysqlComentarios) GetByID(ctx context.Context, IDComento int64, IDComentado int64) (*models.Comentarios, error) {
	query := "Select IDComento, IDComentado, Comentario, Fecha, Hora From Comentarios where IDComento=? AND IDComentado=?"

	rows, err := m.fetch(ctx, query, IDComento, IDComentado)
	if err != nil {
		return nil, err
	}

	payload := &models.Comentarios{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, models.ErrNotFound
	}

	return payload, nil
}

func (m *mysqlComentarios) GetAllByID(ctx context.Context, IDComentado int64) ([]*models.Comentarios, error) {
	query := "Select IDComento, IDComentado, Comentario, Fecha, Hora  From Comentarios where IDComentado=?"

	return m.fetch(ctx, query, IDComentado)
}

func (m *mysqlComentarios) Create(ctx context.Context, c *models.Comentarios) (int64, error) {
	query := "Insert Comentarios SET IDComento=?, IDComentado=?, Comentario=?, Fecha=?, Hora=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	res, err := stmt.ExecContext(ctx, c.IDComento, c.IDComentado, c.Comentario, c.Fecha, c.Hora)
	defer stmt.Close()

	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func (m *mysqlComentarios) Update(ctx context.Context, c *models.Comentarios) (*models.Comentarios, error) {
	query := "Update Comentarios SET Comentario = ?, Fecha = ?, Hora = ? where IDComento=? AND IDComentado=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		c.Comentario,
		c.Fecha,
		c.Hora,
		c.IDComento,
		c.IDComentado,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return c, nil
}

func (m *mysqlComentarios) Delete(ctx context.Context, IDComento int64, IDComentado int64) (bool, error) {
	query := "Delete From Comentarios Where IDComento=? AND IDComentado=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}
	_, err = stmt.ExecContext(ctx, IDComento, IDComentado)
	if err != nil {
		return false, err
	}
	return true, nil
}
