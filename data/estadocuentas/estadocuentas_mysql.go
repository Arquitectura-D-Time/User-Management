package estadocuentas

import (
	"context"
	"database/sql"

	repo "project_user-management_ms/data"
	models "project_user-management_ms/models"
)

func NewSQLEstadoCuentas(Conn *sql.DB) repo.EstadoCuentas {
	return &mysqlEstadoCuentas{
		Conn: Conn,
	}
}

type mysqlEstadoCuentas struct {
	Conn *sql.DB
}

func (m *mysqlEstadoCuentas) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.EstadoCuentas, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.EstadoCuentas, 0)
	for rows.Next() {
		data := new(models.EstadoCuentas)

		err := rows.Scan(
			&data.ID,
			&data.Estado,
			&data.FechaInicio,
			&data.FechaFinal,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (m *mysqlEstadoCuentas) Fetch(ctx context.Context, num int64) ([]*models.EstadoCuentas, error) {
	query := "Select ID, Estado, FechaInicio, FechaFinal  From EstadoCuentas limit ?"

	return m.fetch(ctx, query, num)
}

func (m *mysqlEstadoCuentas) GetByID(ctx context.Context, ID int64) (*models.EstadoCuentas, error) {
	query := "Select ID, Estado, FechaInicio, FechaFinal From EstadoCuentas where ID=?"

	rows, err := m.fetch(ctx, query, ID)
	if err != nil {
		return nil, err
	}

	payload := &models.EstadoCuentas{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, models.ErrNotFound
	}

	return payload, nil
}

func (m *mysqlEstadoCuentas) Create(ctx context.Context, c *models.EstadoCuentas) (int64, error) {
	query := "Insert EstadoCuentas SET ID=?, Estado=?, FechaInicio=?, FechaFinal=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	res, err := stmt.ExecContext(ctx, c.ID, c.Estado, c.FechaInicio, c.FechaFinal)
	defer stmt.Close()

	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func (m *mysqlEstadoCuentas) Update(ctx context.Context, c *models.EstadoCuentas) (*models.EstadoCuentas, error) {
	query := "Update EstadoCuentas SET Estado = ?, FechaInicio = ?, FechaFinal = ? where ID=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		c.Estado,
		c.FechaInicio,
		c.FechaFinal,
		c.ID,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return c, nil
}

func (m *mysqlEstadoCuentas) Delete(ctx context.Context, ID int64) (bool, error) {
	query := "Delete From EstadoCuentas Where ID=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}
	_, err = stmt.ExecContext(ctx, ID)
	if err != nil {
		return false, err
	}
	return true, nil
}
