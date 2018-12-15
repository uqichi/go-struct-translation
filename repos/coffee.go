package repos

import (
	"context"

	"github.com/pkg/errors"

	"github.com/gobuffalo/uuid"

	"github.com/gobuffalo/pop"

	"github.com/uqichi/goffee-server/models"
)

type CoffeeRepository interface {
	Select(ctx context.Context, id uuid.UUID) (*models.Coffee, error)
	Insert(ctx context.Context, coffee *models.Coffee) error
}

type coffeeRepository struct {
	db *pop.Connection
}

func NewCoffeeRepository(db *pop.Connection) CoffeeRepository {
	return &coffeeRepository{db}
}

func (r *coffeeRepository) Select(ctx context.Context, id uuid.UUID) (*models.Coffee, error) {
	coffee := new(models.Coffee)
	err := r.db.Eager().Find(coffee, id)
	return coffee, err
}

func (r *coffeeRepository) Insert(ctx context.Context, coffee *models.Coffee) error {
	verrs, err := r.db.ValidateAndCreate(coffee)
	if err != nil {
		return err
	}
	if verrs.HasAny() {
		return errors.New(verrs.Error())
	}
	return nil
}
