package drug

import (
	"context"

	"github.com/jeffleon1/ionix/pkg/entities"
)

type Controller interface {
	Create(ctx context.Context, entity *entities.Drug) error
	GetAll(ctx context.Context) (*[]entities.Drug, error)
	Update(ctx context.Context, entity *entities.Drug, id uint) error
	Delete(ctx context.Context, id uint) error
	Get(ctx context.Context, id uint) (*entities.Drug, error)
}

type controller struct {
	repo entities.Repository[entities.Drug]
}

func New(repo entities.Repository[entities.Drug]) Controller {
	return &controller{repo}
}

func (c *controller) Get(ctx context.Context, id uint) (*entities.Drug, error) {
	return c.repo.GetByID(ctx, id)
}

func (c *controller) Create(ctx context.Context, model *entities.Drug) error {
	return c.repo.Create(ctx, model)
}

func (c *controller) Update(ctx context.Context, model *entities.Drug, id uint) error {
	return c.repo.Update(ctx, model, id)
}

func (c *controller) Delete(ctx context.Context, id uint) error {
	return c.repo.Delete(ctx, id)
}

func (c *controller) GetAll(ctx context.Context) (*[]entities.Drug, error) {
	return c.repo.GetAll(ctx)
}
