package vaccination

import (
	"context"

	"github.com/jeffleon1/ionix/pkg/entities"
)

type Controller interface {
	Create(ctx context.Context, entity *entities.Vaccination) error
	GetAll(ctx context.Context) (*[]entities.Vaccination, error)
	Update(ctx context.Context, entity *entities.Vaccination, id uint) error
	Delete(ctx context.Context, id uint) error
	Get(ctx context.Context, id uint) (*entities.Vaccination, error)
}

type controller struct {
	repo entities.Repository[entities.Vaccination]
}

func New(repo entities.Repository[entities.Vaccination]) Controller {
	return &controller{repo}
}

func (c *controller) Get(ctx context.Context, id uint) (*entities.Vaccination, error) {
	return c.repo.GetByID(ctx, id)
}

func (c *controller) Create(ctx context.Context, model *entities.Vaccination) error {
	return c.repo.Create(ctx, model)
}

func (c *controller) Update(ctx context.Context, model *entities.Vaccination, id uint) error {
	return c.repo.Update(ctx, model, id)
}

func (c *controller) Delete(ctx context.Context, id uint) error {
	return c.repo.Delete(ctx, id)
}

func (c *controller) GetAll(ctx context.Context) (*[]entities.Vaccination, error) {
	return c.repo.GetAll(ctx)
}
