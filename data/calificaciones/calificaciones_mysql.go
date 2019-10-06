package calificaciones

import (
	"context"
	"database/sql"

	repo "User-Management/data"
	models "User-Management/models"
)

func NewSQLCalificacion(Conn *sql.DB) repo.Calificaciones {
	return &mysqlCalificaciones{
		Conn: Conn,
	}
}

type mysqlCalificaciones struct {
	Conn *sql.DB
}

func (m *mysqlCalificaciones) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Calificaciones, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.Calificaciones, 0)
	for rows.Next() {
		data := new(models.Calificaciones)

		err := rows.Scan(
			&data.IDCalifico,
			&data.IDCalificado,
			&data.Calificacion,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (m *mysqlCalificaciones) Fetch(ctx context.Context, num int64) ([]*models.Calificaciones, error) {
	query := "Select IDCalifico, IDCalificado, Calificacion From Calificaciones limit ?"

	return m.fetch(ctx, query, num)
}

func (m *mysqlCalificaciones) GetByID(ctx context.Context, IDCalifico int64, IDCalificado int64) (*models.Calificaciones, error) {
	query := "Select IDCalifico, IDCalificado, Calificacion From Calificaciones where IDCalifico=? AND IDCalificado=?"

	rows, err := m.fetch(ctx, query, IDCalifico, IDCalificado)
	if err != nil {
		return nil, err
	}

	payload := &models.Calificaciones{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, models.ErrNotFound
	}

	return payload, nil
}

func (m *mysqlCalificaciones) Create(ctx context.Context, c *models.Calificaciones) (int64, error) {
	query := "Insert Calificaciones SET IDCalifico=?, IDCalificado=?, Calificacion=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	res, err := stmt.ExecContext(ctx, c.IDCalifico, c.IDCalificado, c.Calificacion)
	defer stmt.Close()

	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func (m *mysqlCalificaciones) Update(ctx context.Context, c *models.Calificaciones) (*models.Calificaciones, error) {
	query := "Update Calificaciones SET Calificacion = ? where IDCalifico=? AND IDCalificado=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		c.Calificacion,
		c.IDCalifico,
		c.IDCalificado,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return c, nil
}

func (m *mysqlCalificaciones) Delete(ctx context.Context, IDCalifico int64, IDCalificado int64) (bool, error) {
	query := "Delete From Calificaciones Where IDCalifico=? AND IDCalificado=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}
	_, err = stmt.ExecContext(ctx, IDCalifico, IDCalificado)
	if err != nil {
		return false, err
	}
	return true, nil
}
