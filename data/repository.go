package data

import (
	"context"

	"project_user-management_ms/models"
)

type Comentarios interface {
	Fetch(ctx context.Context, num int64) ([]*models.Comentarios, error)
	GetByID(ctx context.Context, id int64, id2 int64) (*models.Comentarios, error)
	Create(ctx context.Context, p *models.Comentarios) (int64, error)
	Update(ctx context.Context, p *models.Comentarios) (*models.Comentarios, error)
	Delete(ctx context.Context, id int64, id2 int64) (bool, error)
}

type Calificaciones interface {
	Fetch(ctx context.Context, num int64) ([]*models.Calificaciones, error)
	GetByID(ctx context.Context, id int64, id2 int64) (*models.Calificaciones, error)
	Create(ctx context.Context, p *models.Calificaciones) (int64, error)
	Update(ctx context.Context, p *models.Calificaciones) (*models.Calificaciones, error)
	Delete(ctx context.Context, id int64, id2 int64) (bool, error)
}

type EstadoCuentas interface {
	Fetch(ctx context.Context, num int64) ([]*models.EstadoCuentas, error)
	GetByID(ctx context.Context, id int64) (*models.EstadoCuentas, error)
	Create(ctx context.Context, p *models.EstadoCuentas) (int64, error)
	Update(ctx context.Context, p *models.EstadoCuentas) (*models.EstadoCuentas, error)
	Delete(ctx context.Context, id int64) (bool, error)
}
