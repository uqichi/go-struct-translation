package usecases

import (
	"context"

	"github.com/gobuffalo/uuid"

	"github.com/uqichi/goffee-server/repos"

	"github.com/uqichi/goffee-server/models"
)

type CoffeeUseCase interface {
	Find(ctx context.Context, id uuid.UUID) (*models.Coffee, error)
	Create(ctx context.Context, coffee *models.Coffee) (*models.Coffee, error)
}

type coffeeUseCase struct {
	repo repos.CoffeeRepository
}

func NewCoffeeUseCase(repo repos.CoffeeRepository) CoffeeUseCase {
	return &coffeeUseCase{repo}
}

func (uc *coffeeUseCase) Find(ctx context.Context, id uuid.UUID) (*models.Coffee, error) {
	return uc.repo.Select(ctx, id)
}

func (uc *coffeeUseCase) Create(ctx context.Context, coffee *models.Coffee) (*models.Coffee, error) {
	err := uc.repo.Insert(ctx, coffee)
	if err != nil {
		return nil, err
	}
	return uc.repo.Select(ctx, coffee.ID)
}
